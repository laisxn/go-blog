package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin/config"
)

func AuthApiKeyMiddleware(c *gin.Context) {

	post_api_auth_key := c.PostForm("api_auth_key")
	api_auth_key := config.Get("app.api_auth_key")

	if api_auth_key != post_api_auth_key {
		c.Abort()
	}
}
