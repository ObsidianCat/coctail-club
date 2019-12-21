package main

import (
	"cocktail-club/http_handlers"
	"cocktail-club/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.QueryChecker())
	r.Use(middleware.LogWriter())

	r.GET("/ping", http_handlers.Ping)
	r.GET("/cocktail", http_handlers.Cocktail)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
