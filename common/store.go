package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	cocktailsList Cocktails
	ById          map[int]Cocktail
	ByIngredient  map[string][]int
}

//creating new store
func NewStorage() *Store {
	return &Store{}
}

//load cocktails recipes from file
func (s Store) LoadCocktails() {
	// read file
	data, err := ioutil.ReadFile("./data/cocktail_recipes.json")
	if err != nil {
		fmt.Print(err)
	}

	// json data
	cocktailsList := []Cocktail{}
	//var cocktailsList []Cocktail

	err = json.Unmarshal(data, &cocktailsList)
	if err != nil {
		fmt.Println("error:", err)
	}

	s.cocktailsList = cocktailsList
	s.ById = convertCocktailsListIntoById(cocktailsList)
	s.ByIngredient = convertCocktailsListInByIngredient(cocktailsList)
}

func convertCocktailsListIntoById(cocktails []Cocktail) map[int]Cocktail {
	byId := make(map[int]Cocktail)
	for _, c := range cocktails {
		byId[c.Id] = c
	}
	return byId
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
