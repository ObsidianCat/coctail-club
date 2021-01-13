package cocktails

import (
	"cocktail-club/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// GetByName http request handler
func GetByName(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))

	cocktailBytes, _ := common.ProxyRequest(common.CocktailDbURLSearchByName + name)
	result, err := common.TransformAPIBytesToCtails(cocktailBytes)

	if err != nil {
		log.Println(err.Error())
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(http.StatusNotFound, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// GetByIngredient http request handler
func GetByIngredient(c *gin.Context) {
	ingr := strings.ToLower(c.Param("ingredient"))

	cocktailBytes, _ := common.ProxyRequest(common.CocktailDbURLSeacrhByIngredient + ingr)
	result, err := common.TransformAPIBytesToCtailPreview(cocktailBytes)
	if err != nil {
		log.Println(err.Error())
		c.Header(common.ErrorHeaderName, "Ingredient not found")
		c.JSON(http.StatusNotFound, gin.H{
			common.ErrorMessageKey: "No cocktails with this ingredient",
		})
	} else {
		c.JSON(http.StatusOK, result)

	}
}

// GetByID http request handler
func GetByID(c *gin.Context) {
	id := strings.ToLower(c.Param("id"))

	cocktailBytes, _ := common.ProxyRequest(common.CocktailDbURLLookupByID + id)
	result, err := common.TransformAPIBytesToCtails(cocktailBytes)

	if err != nil {
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(http.StatusNotFound, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})
		log.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, result)
	}

}
