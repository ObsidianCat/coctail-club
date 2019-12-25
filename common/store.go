package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

// define data structure

//Cocktail item structure
type Cocktail struct {
	Name        string
	Ingredients []string
	ID          int
	Preparation string
	URL         string
}

type Cocktails []Cocktail

type Store struct {
	ByID         map[int]Cocktail
	ByIngredient map[string][]int
	ByName       map[string]int
}

var storePointer *Store

//creating new store
func StoreInit() *Store {
	storePointer = &Store{}
	storePointer.LoadCocktails()
	return storePointer
}

func GetStore() *Store {
	if storePointer == nil {
		StoreInit()
	}
	return storePointer
}

func ReadDataFile() ([]byte, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("cannot read file")
	}

	prefixPath := filepath.Dir(file)
	path := prefixPath + "/cocktail_recipes.json"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return nil, errors.New("cannot read file")
	}

	return data, nil
}

//load cocktails recipes from file into store
func (s *Store) LoadCocktails() {
	// read file
	data, _ := ReadDataFile()

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
