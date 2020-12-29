package collection

import (
	"cocktail-club/collection"
	"cocktail-club/common"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Add http request handler
func Add(c *gin.Context) {
	storeRef := collection.GetStore()

	id := strings.ToLower(c.Param("id"))

	cocktailBytes, _ := common.ProxyRequest(common.CocktailDbURLLookupByID + id)
	result, err := common.TransformAPIBytesToCtails(cocktailBytes)

	if err != nil {
		log.Println(err.Error())

		c.Header(common.ErrorHeaderName, "Cocktail not found")
		c.JSON(http.StatusNotFound, gin.H{
			common.ErrorMessageKey: "Cocktail with this name not found",
		})

	} else {
		storeRef.Cocktails = append(storeRef.Cocktails, result[0])
		c.JSON(http.StatusOK, gin.H{
			"Cocktail Name": result[0].Name,
			"Cocktail Id":   result[0].ID,
		})
	}
}

// Get http request handler
func Get(c *gin.Context) {
	storeRef := collection.GetStore()

	if storeRef != nil {
		c.JSON(http.StatusOK, storeRef.Cocktails)
	} else {
		c.Header(common.ErrorHeaderName, "Cocktail collection not found")
		c.JSON(http.StatusNotFound, gin.H{
			common.ErrorMessageKey: "Cocktail collection not found",
		})
	}
}

// Save http request handler
func Save(c *gin.Context) {
	storeRef := collection.GetStore()

	file, _ := json.MarshalIndent(storeRef.Cocktails, "", " ")

	err := ioutil.WriteFile(common.CocktailCollectionFilePath, file, 0644)

	if err != nil {
		log.Println(err.Error())
		c.Header(common.ErrorHeaderName, "Cannot save collection file")
		c.JSON(http.StatusNotFound, gin.H{
			common.ErrorMessageKey: "Cannot save collection",
		})
	} else {
		c.String(http.StatusOK, "Collections was saved successfully")
	}
}
