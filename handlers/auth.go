package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nqvinh00/CakeAssignment/model"
	"golang.org/x/crypto/bcrypt"
)

func (h *httpd) Login(c *gin.Context) {
	req, ctx := &model.LoginReq{}, c.Request.Context()

	if err := c.ShouldBindJSON(req); err != nil {
		responseJSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if errMessage := req.Valid(); errMessage != model.Success {
		responseJSON(c, http.StatusBadRequest, errMessage, nil)
		return
	}

	user, err := h.userDAO.SelectToLogin(ctx, req.Username)
	if err != nil {
		responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	if user == nil {
		responseJSON(c, http.StatusUnauthorized, "User not found", nil)
		return
	}

	// user and user_sec should have check sum and compare with hashsip

	if user.Status != model.UserActivatedStatus {
		responseJSON(c, http.StatusUnauthorized, "This account has been deactived. Please contact adminstrator", nil)
		return
	}

	userSec, err := h.userSecDAO.Select(ctx, user.ID)
	if err != nil {
		responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	if userSec == nil {
		responseJSON(c, http.StatusUnauthorized, model.UserNotFound, nil)
		return
	}

	if err = bcrypt.CompareHashAndPassword(userSec.Password, []byte(req.Password)); err != nil {
		user.LoginAttempt--
		if user.LoginAttempt <= 0 {
			user.Status = model.UserDeactivatedStatus
		}

		// skip error
		_ = h.userDAO.Update(ctx, user)
		responseJSON(c, http.StatusUnauthorized, fmt.Sprintf("Password is incorrect. You only have %d times left to retry", user.LoginAttempt), nil)
		return
	}

	// restart attempts to 3 just in case user retry
	// TODO: use const or define config
	user.LoginAttempt = 3
	_ = h.userDAO.Update(ctx, user)

	// TODO: define config
	expire := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claim{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
		},
	})

	t, err := token.SignedString([]byte(h.config.SecretKey))
	if err != nil {
		responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	// TODO: error check
	_ = h.userDAO.LastLogin(ctx, user.ID)

	responseJSON(c, http.StatusOK, "Login success", gin.H{"token": t})
}

func (h *httpd) Signup(c *gin.Context) {
	req, ctx := &model.NewUserReq{}, c.Request.Context()

	if err := c.ShouldBindJSON(req); err != nil {
		responseJSON(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if errMessage := req.Valid(); errMessage != model.Success {
		responseJSON(c, http.StatusBadRequest, errMessage, nil)
		return
	}
	userAccount := &model.User{
		Fullname:    req.Fullname,
		Email:       req.Email,
		Username:    req.Username,
		PhoneNumber: req.PhoneNumber,
		// CampaignID: ,
		Birthday: req.Birthday,
	}

	user, err := h.userDAO.SelectToSignup(ctx, req.Username, req.Email, req.PhoneNumber)
	if err != nil {
		responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	if user != nil {
		responseJSON(c, http.StatusBadRequest, model.UserAlreadyExists, nil)
		return
	}

	id, err := h.userDAO.Insert(ctx, userAccount)
	if err != nil {
		responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	p, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	if err := h.userSecDAO.Insert(ctx, &model.UserSec{
		UserID:   id,
		Password: p,
	}); err != nil {
		responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	responseJSON(c, http.StatusOK, "Signup success", nil)
}
