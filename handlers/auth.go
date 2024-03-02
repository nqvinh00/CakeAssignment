package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nqvinh00/CakeAssignment/model"
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

	token, err := h.authenticator.Login(ctx, req.Username, req.Password)
	if err != nil {
		if err == model.ErrUserNotFound {
			responseJSON(c, http.StatusUnauthorized, err.Error(), nil)
		} else {
			responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		}
		
		return
	}

	responseJSON(c, http.StatusOK, "Login success", gin.H{"token": token})
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

	user := &model.User{
		Fullname:    req.Fullname,
		Email:       req.Email,
		Username:    req.Username,
		PhoneNumber: req.PhoneNumber,
		// CampaignID: ,
		Birthday: req.Birthday,
	}

	if err := h.authenticator.CreateUser(ctx, user, req.Password); err != nil {
		if err == model.ErrUserAlreadyExists {
			responseJSON(c, http.StatusUnauthorized, err.Error(), nil)
		} else {
			responseJSON(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		}

		return
	}

	responseJSON(c, http.StatusOK, "Signup success", nil)
}
