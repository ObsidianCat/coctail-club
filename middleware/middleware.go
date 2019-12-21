package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func QueryChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In Checker middleware")
		c.Next()
	}
}

func LogWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		log.Println("In LogWriter middleware")

		c.Next()
	}
}
