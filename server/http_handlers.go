package server

import (
	"cocktail-club/common"
	"cocktail-club/store"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// Ping http request handler for default test route
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// CocktailByIngredient http request handler
func CocktailByIngredient(c *gin.Context) {
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

// CocktailByName http request handler
func CocktailByName(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))
	store := store.GetStore()
	cID, isFound := store.ByName[name]
	if isFound {
		cocktail := store.ByID[cID]
		c.JSON(200, cocktail)
	} else {
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})
	}

}
