package http_handlers

import (
	"github.com/gin-gonic/gin"
	"log"
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

	c.JSON(200, gin.H{
		"message": "name",
	})
}
