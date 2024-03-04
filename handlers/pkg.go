package handlers

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nqvinh00/CakeAssignment/model"
)

func responseJSON(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, &response{
		Error: responseError{
			Code:    status,
			Message: message,
		},
		Data: data,
	})
}

type responseError struct {
	Code    int
	Message string
}

type response struct {
	Error responseError
	Data  interface{}
}

func (h *httpd) validateToken(signedToken string) (claims *model.Claim, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &model.Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(h.jwtSecretKey), nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*model.Claim)
	if !ok {
		err = errors.New("invalid token")
		return
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		err = errors.New("token expired")
		return
	}

	return
}
