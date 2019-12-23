package main

import (
	"cocktail-club/http_handlers"
	"cocktail-club/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", http_handlers.Ping)
	r.GET("/cocktail/ingredient/", middleware.QueryChecker(), http_handlers.CocktailByIngredient, middleware.LogWriter())
	r.GET("/cocktail/name/:name", middleware.QueryChecker(), http_handlers.CocktailByName, middleware.LogWriter())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
