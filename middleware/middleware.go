package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func QueryChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()
		if len(query) <= 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "ingredient for search as Query parameter is expected"})
		}
	}
}

func LogWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		log.Println("In LogWriter middleware")

		c.Next()
	}
}
