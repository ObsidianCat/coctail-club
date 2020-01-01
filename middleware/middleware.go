package middleware

import (
	"cocktail-club/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// QueryChecker validate that search query is nto empty and prevent request form going further if it is.
func QueryChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()
		if len(query) <= 0 {
			c.Header(common.ErrorHeaderName, "Query parameter is required")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{common.ErrorMessageKey: "Expect ingredient as search query parameter"})
		}
	}
}
