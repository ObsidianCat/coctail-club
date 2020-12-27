package cocktails

import (
	"cocktail-club/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// GetByName http request handler
func GetByName(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))

	cocktailBytes := common.ProxyRequest(common.CocktailDbUrlSearchByName + name)
	result := common.TransformApiBytesToCtails(cocktailBytes)

	if result != nil {
		c.JSON(200, result)
	} else {
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})
	}
}

// GetByIngredient http request handler
func GetByIngredient(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))

	cocktailBytes := common.ProxyRequest(common.CocktailDbUrlSeacrhByIngredient + name)
	result := common.TransformApiBytesToCtailPreview(cocktailBytes)

	if result != nil {
		c.JSON(200, result)
	} else {
		c.Header(common.ErrorHeaderName, "Ingredient not found")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "No cocktails with this ingredient",
		})
	}
}

// GetById http request handler
func GetById(c *gin.Context) {
	id := strings.ToLower(c.Param("id"))

	cocktailBytes := common.ProxyRequest(common.CocktailDbUrlLookupById + id)
	result := common.TransformApiBytesToCtails(cocktailBytes)

	if result != nil {
		c.JSON(200, result)
	} else {
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})
	}

}
