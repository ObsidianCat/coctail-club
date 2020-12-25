package cocktain_by_id

import (
	"cocktail-club/common"
	"cocktail-club/http_handlers"
	"github.com/gin-gonic/gin"
	"strings"
)

// CocktailById http request handler
func CocktailById(c *gin.Context) {
	id := strings.ToLower(c.Param("id"))

	cocktailBytes := http_handlers.ProxyRequest(common.COCKTAIL_DB_URL_LOOKUP_BY_ID + id)
	result := http_handlers.TransformApiBytesToCtails(cocktailBytes)

	if result != nil {
		c.JSON(200, result)
	} else {
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})
	}

}
