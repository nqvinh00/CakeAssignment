package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nqvinh00/CakeAssignment/dao"
)

type httpd struct {
	userDAO    dao.IUserDAO
	userSecDAO dao.IUserSecDAO
}

func NewHTTPD(userDAO dao.IUserDAO, userSecDAO dao.IUserSecDAO) *httpd {
	return &httpd{
		userDAO:    userDAO,
		userSecDAO: userSecDAO,
	}
}

func (h *httpd) SetupRouter() *gin.Engine {
	r := gin.Default()

	// some middlewares here

	// auth group for authentication
	auth := r.Group("/auth")
	auth.POST("/login")
	auth.POST("/signup")

	//
	return r
}

func responseJSON(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(http.StatusOK, &response{
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
