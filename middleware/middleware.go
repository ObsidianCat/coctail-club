package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func QueryChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In Checker middleware 1")

		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
		log.Println("In Checker middleware 2")

	}
}

func LogWriter() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before request
		log.Println("In LogWriter middleware 1")

		c.Next()

		// access the status we are sending
		log.Println("In LogWriter middleware 2")
	}
}
