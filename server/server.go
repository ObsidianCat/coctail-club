package server

import (
	"cocktail-club/httphandlers"
	"cocktail-club/httphandlers/cocktails"
	"cocktail-club/httphandlers/collection"
	"github.com/gin-gonic/gin"
)

// SetupRouter initialize Gin Engine, define routes and return router reference
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", httphandlers.Ping)

	router.GET("/cocktails/ingredient/:ingredient", cocktails.GetByIngredient)
	router.GET("/cocktails/name/:name", cocktails.GetByName)
	router.GET("/cocktails/id/:id", cocktails.GetByID)

	router.PUT("/collection/add/id/:id", collection.Add)
	router.POST("/collection", collection.Save)
	router.GET("/collection", collection.Get)

	return router
}
