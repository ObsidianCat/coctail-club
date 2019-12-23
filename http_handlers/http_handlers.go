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

func CocktailByIngredient(c *gin.Context) {
	log.Println("In cocktail main handler")

	c.JSON(200, gin.H{
		"message": "ingredient",
	})
}

func CocktailByName(c *gin.Context) {
	log.Println("In cocktail main handler")
	name := strings.ToLower(c.Param("name"))
	store := common.GetStore()
	cId, isFound := store.ByName[name]
	if isFound {
		cocktail := store.ById[cId]
		c.JSON(200, cocktail)
	} else {
		c.JSON(404, gin.H{
			"cocktail": name,
		})
	}

}
