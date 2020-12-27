package common

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"runtime"

	//"cocktail-club/collection"
	"fmt"
	"github.com/oleiade/reflections"
	"io/ioutil"
	"net/http"
	"strconv"
)

type apiSearchResultsByName struct {
	Cocktails []ApiCocktail `json:"drinks"`
}
type apiSearchResultsByIngredient struct {
	Cocktails []CocktailPreview `json:"drinks"`
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

func intToString(num int) string {
	result := strconv.Itoa(num)
	return result
}

func ApiCtailToCtail(input ApiCocktail) Cocktail {
	var ingedientsList []string
	for i := 1; i <= 5; i++ {
		value, _ := reflections.GetField(input, "StrIngredient"+intToString(i))
		if value != "" {
			ingedientsList = append(ingedientsList, value.(string))
		}
	}
	cTail := Cocktail{
		Name:        input.Name,
		Ingredients: ingedientsList,
		ID:          stringToInt(input.ID),
		Preparation: input.Preparations,
		Image:       input.Image,
	}
	return cTail
}

func TransformApiBytesToCtails(cocktailBytes []byte) []Cocktail {
	var cocktailsListFromApi apiSearchResultsByName
	var cocktailsResults []Cocktail

	err := json.Unmarshal(cocktailBytes, &cocktailsListFromApi)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, apiCTail := range cocktailsListFromApi.Cocktails {
		cocktailsResults = append(cocktailsResults, ApiCtailToCtail(apiCTail))
	}
	return cocktailsResults
}

func TransformApiBytesToCtailPreview(cocktailBytes []byte) []CocktailPreview {
	var cocktailsListFromApi apiSearchResultsByIngredient
	//var cocktailsResults []CocktailPreview

	err := json.Unmarshal(cocktailBytes, &cocktailsListFromApi)
	if err != nil {
		fmt.Println("error:", err)
	}

	return cocktailsListFromApi.Cocktails
}

func ReadDataFileWithPathFromCallerFile(fileName string) ([]byte, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("cannot read file")
	}

	prefixPath := filepath.Dir(file)
	path := prefixPath + "/" + fileName

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return nil, errors.New("cannot read file")
	}

	return data, nil
}
