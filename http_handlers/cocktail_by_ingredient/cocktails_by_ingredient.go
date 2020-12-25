package cocktail_by_ingredient

import (
	"cocktail-club/common"
	"cocktail-club/store"
	"github.com/gin-gonic/gin"
	"log"
)

// CocktailsByIngredient http request handler
func CocktailsByIngredient(c *gin.Context) {
	storeRef := store.GetStore()

	var queryKeys []string
	for k := range c.Request.URL.Query() {
		queryKeys = append(queryKeys, k)
	}

	mainKey := queryKeys[0]

	foundCocktailsIds := storeRef.ByIngredient[mainKey]
	log.Println("Found cocktail ids")
	log.Println(foundCocktailsIds)

	if len(foundCocktailsIds) <= 0 {
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(404, gin.H{common.ErrorMessageKey: "No results for this ingredient"})
		return
	}

	foundCocktails := storeRef.FindCocktailsByIds(foundCocktailsIds)
	c.JSON(200, foundCocktails)
}
