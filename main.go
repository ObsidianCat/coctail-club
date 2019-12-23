package main

import (
	"cocktail-club/common"
	"cocktail-club/http_handlers"
	"cocktail-club/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := common.NewStorage()
	store.LoadCocktails()

	r.GET("/ping", http_handlers.Ping)
	r.GET("/cocktail/ingredient/", middleware.QueryChecker(), http_handlers.CocktailByIngredient, middleware.LogWriter())
	r.GET("/cocktail/name/", middleware.QueryChecker(), http_handlers.CocktailByName, middleware.LogWriter())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
