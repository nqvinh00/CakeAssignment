package handlers

import (
	"net/http"

	_ "github.com/nqvinh00/CakeAssignment/handlers/docs"
	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/nqvinh00/CakeAssignment/services"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type httpd struct {
	config        model.HTTP
	jwtSecretKey  string
	authenticator services.Authenticator
}

func NewHTTPD(config model.HTTP, authenticator services.Authenticator, jwtSecretKey string) *httpd {
	return &httpd{
		config:        config,
		jwtSecretKey:  jwtSecretKey,
		authenticator: authenticator,
	}
}

// @title           Cake Interview Assignment
// @version         1.0

// @contact.email  nqvinh00@gmail.com

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

func (h *httpd) SetupRouter() *gin.Engine {
	r := gin.Default()

	// some middlewares here

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// auth group for authentication
	auth := r.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/signup", h.Signup)

	r.Use(h.authMiddleware())
	r.Any("/ping", Ping) // Test authentication
	return r
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
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
