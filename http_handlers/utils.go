package http_handlers

import (
	"cocktail-club/store"
	"encoding/json"
	//"cocktail-club/store"
	"fmt"
	"github.com/oleiade/reflections"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ApiCocktail struct {
	Name           string `json:"strDrink"`
	ID             string `json:"idDrink"`
	StrIngredient1 string `json:"strIngredient1"`
	StrIngredient2 string `json:"strIngredient2"`
	StrIngredient3 string `json:"strIngredient3"`
	StrIngredient4 string `json:"strIngredient4"`
	StrIngredient5 string `json:"strIngredient5"`
	strMeasure1    string `json:"strMeasure1"`
	strMeasure2    string `json:"strMeasure2"`
	StrMeasure3    string `json:"strMeasure3"`
	StrMeasure4    string `json:"strMeasure4"`
	StrMeasure5    string `json:"strMeasure5"`
	Preparations   string `json:"strInstructions"`
}

type ApiSearchResults struct {
	Cocktails []ApiCocktail `json:"drinks"`
}

// ProxyRequest proxies request to external cocktail API
func ProxyRequest(cocktailsApiUrl string) []byte {
	res, _ := http.Get(cocktailsApiUrl)
	var bytes []byte
	bytes, _ = ioutil.ReadAll(res.Body)
	return bytes
}

func stringToInt(st string) int {
	result, _ := strconv.ParseInt(st, 10, 0)
	return int(result)
}

func IntToString(num int) string {
	result := strconv.Itoa(num)
	return result
}

func ApiCtailToCtail(input ApiCocktail) store.Cocktail {
	var ingedientsList []string
	for i := 1; i <= 5; i++ {
		value, _ := reflections.GetField(input, "StrIngredient"+IntToString(i))
		if value != "" {
			ingedientsList = append(ingedientsList, value.(string))
		}
	}
	cTail := store.Cocktail{
		Name:        input.Name,
		Ingredients: ingedientsList,
		ID:          stringToInt(input.ID),
		Preparation: input.Preparations,
		URL:         "",
	}
	return cTail
}
func TransformApiBytesToCtails(cocktailBytes []byte) []store.Cocktail {
	var cocktailsListFromApi ApiSearchResults
	var cocktailsResults []store.Cocktail

	err := json.Unmarshal(cocktailBytes, &cocktailsListFromApi)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, apiCTail := range cocktailsListFromApi.Cocktails {
		cocktailsResults = append(cocktailsResults, ApiCtailToCtail(apiCTail))
	}
	return cocktailsResults
}
