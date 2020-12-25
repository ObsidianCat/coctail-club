package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

//Cocktail item is a main data unit for the store
type Cocktail struct {
	Name        string
	Ingredients []string
	ID          int
	Preparation string
	URL         string
}

// Cocktails is a list of Cocktail type items
type Cocktails []Cocktail

// Store contains all data for the service
type Store struct {
	ByID         map[int]Cocktail
	ByIngredient map[string][]int
	ByName       map[string]int
	dataPath     string
}

var storePointer *Store

//StoreInit create new store instance in memory and return reference to the newly created store
func StoreInit(params ...string) *Store {
	path := "cocktail_recipes.json"
	if len(params) > 0 {
		path = params[0]
	}
	storePointer = &Store{dataPath: path}
	storePointer.LoadCocktails()
	return storePointer
}

// GetStore return reference to existing store or create new store and return its reference
func GetStore() *Store {
	if storePointer == nil {
		StoreInit()
	}
	return storePointer
}

func ReadDataFile(fileName string) ([]byte, error) {
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

// LoadCocktails loads cocktail recipes from file into store
func (s *Store) LoadCocktails() {
	// read file
	data, _ := ReadDataFile(s.dataPath)

	// json data
	cocktailsList := []Cocktail{}
	//var cocktailsList []Cocktail

	err := json.Unmarshal(data, &cocktailsList)
	if err != nil {
		fmt.Println("error:", err)
	}

	//(*s) converting pointer of the receiver to the value of the receiver
	(*s).ByID = convertCocktailsListIntoByID(cocktailsList)
	(*s).ByIngredient = convertCocktailsListInByIngredient(cocktailsList)
	(*s).ByName = convertCocktailsListIntoByName(cocktailsList)
}

func convertCocktailsListIntoByID(cocktails []Cocktail) map[int]Cocktail {
	byID := make(map[int]Cocktail)
	for _, c := range cocktails {
		byID[c.ID] = c
	}
	return byID
}
func convertCocktailsListIntoByName(cocktails []Cocktail) map[string]int {
	byName := make(map[string]int)
	for _, c := range cocktails {
		byName[strings.ToLower(c.Name)] = c.ID
	}
	return byName
}
func convertCocktailsListInByIngredient(cocktails []Cocktail) map[string][]int {
	byIngredient := make(map[string][]int)
	for _, c := range cocktails {
		//for ingredients list in cocktails
		for _, ingredient := range c.Ingredients {
			//split into separate words
			searchTerms := strings.Fields(ingredient)

			//for every search word
			for _, term := range searchTerms {
				termNormalized := strings.ToLower(term)
				elem, ok := byIngredient[termNormalized]
				if ok {
					byIngredient[termNormalized] = append(elem, c.ID)
				} else {
					byIngredient[termNormalized] = []int{c.ID}
				}
			}
		}
	}
	return byIngredient
}

// FindCocktailsByIds accepts cocktail ids as list of ints and return list of cocktail objects
func (s *Store) FindCocktailsByIds(ids []int) []Cocktail {
	var foundCocktails Cocktails
	for _, id := range ids {
		foundCocktails = append(foundCocktails, storePointer.ByID[id])
	}

	return foundCocktails
}
