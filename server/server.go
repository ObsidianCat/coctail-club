package server

import (
	"cocktail-club/http_handlers"
	"cocktail-club/http_handlers/cocktail_by_ingredient"
	"cocktail-club/http_handlers/cocktail_by_name"
	"cocktail-club/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter initialize Gin Engine, define routes and return router reference
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", http_handlers.Ping)
	r.GET("/cocktail/ingredient", middleware.QueryChecker(), cocktail_by_ingredient.CocktailsByIngredient)
	r.GET("/cocktail/name/:name", cocktail_by_name.CocktailByName)

	return r
}
