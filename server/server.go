package server

import (
	"cocktail-club/http_handlers"
	"cocktail-club/http_handlers/cocktails"
	"cocktail-club/http_handlers/collection"
	"github.com/gin-gonic/gin"
)

// SetupRouter initialize Gin Engine, define routes and return router reference
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", http_handlers.Ping)

	router.GET("/cocktails/ingredient/:name", cocktails.GetByIngredient)
	router.GET("/cocktails/name/:name", cocktails.GetByName)
	router.GET("/cocktails/id/:id", cocktails.GetById)

	router.PUT("/collection/add/id/:id", collection.Add)
	router.POST("/collection", collection.Save)
	router.GET("/collection", collection.Get)

	return router
}
