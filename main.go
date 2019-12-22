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
	r.GET("/cocktail", middleware.QueryChecker(), http_handlers.Cocktail, middleware.LogWriter())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
