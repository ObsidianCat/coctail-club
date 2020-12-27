package httphandlers

import "github.com/gin-gonic/gin"

// Ping http request handler for default test route
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
