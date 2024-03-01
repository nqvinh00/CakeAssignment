package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nqvinh00/CakeAssignment/dao"
	"github.com/nqvinh00/CakeAssignment/model"
)

type httpd struct {
	config     model.HTTP
	userDAO    dao.IUserDAO
	userSecDAO dao.IUserSecDAO
}

func NewHTTPD(config model.HTTP, userDAO dao.IUserDAO, userSecDAO dao.IUserSecDAO) *httpd {
	return &httpd{
		config:     config,
		userDAO:    userDAO,
		userSecDAO: userSecDAO,
	}
}

func (h *httpd) SetupRouter() *gin.Engine {
	r := gin.Default()

	// some middlewares here

	// auth group for authentication
	r.Any("/ping", Ping)
	auth := r.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/signup")

	//
	return r
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
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
