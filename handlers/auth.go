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

	if errCode := req.Valid(); errCode != model.Success {
		responseJSON(c, int(errCode), model.ErrorMessage[errCode], nil)
		return
	}

	// temp
	user, err := h.userDAO.Select(ctx, req.Username)
	if err != nil || user == nil {
		responseJSON(c, http.StatusUnauthorized, "User not found", nil)
		return
	}

	// user and user_sec should have check sum and compare with hashsip

	if user.Status != model.UserActivatedStatus {
		responseJSON(c, http.StatusUnauthorized, "This account has been deactived. Please contact adminstrator", nil)
		return
	}

	userSec, err := h.userSecDAO.Select(ctx, user.ID)
	if err != nil || userSec == nil {
		responseJSON(c, http.StatusUnauthorized, "User not found", nil)
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, nil)
}
