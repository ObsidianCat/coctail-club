package http_handlers

import (
	"cocktail-club/common"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// CocktailByIngredient http request handler
func CocktailByIngredient(c *gin.Context) {
	store := common.GetStore()
	log.Println("In cocktail by ingridient main handler")

	var queryKeys []string
	for k := range c.Request.URL.Query() {
		queryKeys = append(queryKeys, k)
	}

	mainKey := queryKeys[0]

	foundCoctailsIds := store.ByIngredient[mainKey]
	var foundCoctails common.Cocktails
	for _, id := range foundCoctailsIds {
		foundCoctails = append(foundCoctails, store.ByID[id])
	}

	c.JSON(200, foundCoctails)
}

// CocktailByName http request handler
func CocktailByName(c *gin.Context) {
	log.Println("In cocktail main handler")
	name := strings.ToLower(c.Param("name"))
	store := common.GetStore()
	cID, isFound := store.ByName[name]
	if isFound {
		cocktail := store.ByID[cID]
		c.JSON(200, cocktail)
	} else {
		c.JSON(404, gin.H{
			"cocktail": name,
		})
	}

}
