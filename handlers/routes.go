package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/nqvinh00/CakeAssignment/services"
	"github.com/rs/zerolog/log"
)

type httpd struct {
	config        model.HTTP
	jwtSecretKey  string
	authenticator services.Authenticator
}

func NewHTTPD(config model.HTTP, authenticator services.Authenticator, jwtSecretKey string) *httpd {
	return &httpd{
		config:        config,
		jwtSecretKey: jwtSecretKey,
		authenticator: authenticator,
	}
}

func (h *httpd) SetupRouter() *gin.Engine {
	r := gin.Default()

	// some middlewares here

	// auth group for authentication
	r.Any("/ping", Ping)
	auth := r.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/signup", h.Signup)

	r.Use(h.authMiddleware())
	r.Any("/pong", Pong)
	return r
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func Pong(c *gin.Context) {
	log.Info().Str("username", c.GetString("username")).Str("email", c.GetString("email")).Msg("")
	c.JSON(http.StatusOK, gin.H{"message": "ping"})
}

func (h *httpd) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			responseJSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
			c.Abort()
			return
		}

		claim, err := h.validateToken(token)
		if err != nil {
			responseJSON(c, http.StatusInternalServerError, err.Error(), nil)
			c.Abort()
			return
		}
		c.Set("username", claim.Username)
		c.Set("email", claim.Email)
		c.Next()
	}
}
