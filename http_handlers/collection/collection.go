package collection

import (
	"cocktail-club/common"
	"cocktail-club/http_handlers"
	"cocktail-club/store"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// AddToCollection http request handler
func AddToCollection(c *gin.Context) {
	storeRef := store.GetStore()

	id := strings.ToLower(c.Param("id"))

	cocktailBytes := http_handlers.ProxyRequest(common.COCKTAIL_DB_URL_LOOKUP_BY_ID + id)
	result := http_handlers.TransformApiBytesToCtails(cocktailBytes)

	if result != nil {
		storeRef.Cocktails = append(storeRef.Cocktails, result[0])
		c.JSON(200, gin.H{
			"Cocktail Name": result[0].Name,
			"Cocktail Id":   result[0].ID,
		})
	} else {
		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})
	}
}

// AddToCollection http request handler
func Get(c *gin.Context) {
	storeRef := store.GetStore()

	if storeRef != nil {
		c.JSON(200, storeRef.Cocktails)
	} else {
		c.Header(common.ErrorHeaderName, "Cocktail collection not found")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "Cocktail collection not found",
		})
	}
}

// SaveCollection http request handler
func SaveCollection(c *gin.Context) {
	storeRef := store.GetStore()

	file, _ := json.MarshalIndent(storeRef.Cocktails, "", " ")

	err := ioutil.WriteFile(common.CocktailCollectionName, file, 0644)

	if err != nil {
		c.Header(common.ErrorHeaderName, "Cannot save collection file")
		c.JSON(404, gin.H{
			common.ErrorMessageKey: "Cannot save collection",
		})
	} else {
		c.String(http.StatusOK, "Collections was saved successfully")
	}
}
