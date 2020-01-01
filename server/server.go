package server

import (
	"cocktail-club/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter initialize Gin Engine, define routes and return router reference
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.GET("/cocktail/ingredient", middleware.QueryChecker(), CocktailByIngredient)
	r.GET("/cocktail/name/:name", CocktailByName)

	return r
}
