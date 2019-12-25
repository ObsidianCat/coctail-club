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
type Cocktail struct {
	Name        string
	Ingredients []string
	Id          int
	Preparation string
	Url         string
}

type Cocktails []Cocktail

type Store struct {
	ById         map[int]Cocktail
	ByIngredient map[string][]int
	ByName       map[string]int
}

var storePointer *Store

//creating new store
func StoreInit() *Store {
	s := Store{}
	storePointer = &s
	storePointer.LoadCocktails()
	return storePointer
}

func GetStore() *Store {
	if storePointer != nil {
		return storePointer
	} else {
		StoreInit()
		return storePointer
	}
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

//load cocktails recipes from file
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
	(*s).ById = convertCocktailsListIntoById(cocktailsList)
	(*s).ByIngredient = convertCocktailsListInByIngredient(cocktailsList)
	(*s).ByName = convertCocktailsListIntoByName(cocktailsList)
}

func convertCocktailsListIntoById(cocktails []Cocktail) map[int]Cocktail {
	byId := make(map[int]Cocktail)
	for _, c := range cocktails {
		byId[c.Id] = c
	}
	return byId
}
func convertCocktailsListIntoByName(cocktails []Cocktail) map[string]int {
	byName := make(map[string]int)
	for _, c := range cocktails {
		byName[strings.ToLower(c.Name)] = c.Id
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
					byIngredient[termNormalized] = append(elem, c.Id)
				} else {
					byIngredient[termNormalized] = []int{c.Id}
				}
			}
		}
	}
	return byIngredient
}
