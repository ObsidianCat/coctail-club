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

func Cocktail(c *gin.Context) {
	log.Println("In middleware")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
