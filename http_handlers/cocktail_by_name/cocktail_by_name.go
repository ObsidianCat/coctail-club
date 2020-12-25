package cocktail_by_name

import (
	"cocktail-club/common"
	"cocktail-club/http_handlers"
	"github.com/gin-gonic/gin"
	"strings"
)

// CocktailByName http request handler
func CocktailByName(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))

	cocktailBytes := http_handlers.ProxyRequest(common.Cocktail_DB_URL_SEARCH_BY_NAME + name)
	results := http_handlers.TransformApiBytesToCtails(cocktailBytes)

	if results != nil {
		c.JSON(200, results)
	} else {
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})
	}

}
