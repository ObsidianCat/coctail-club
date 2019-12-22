package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// define data structure
type Cocktail struct {
	Name        string
	Ingredients []string
	Id          int
}

type Cocktails []Cocktail

type Store struct {
	cocktailsList Cocktails
	ById          map[int]Cocktail
	ByIngredients map[string][]int
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

	// unmarshall it
	err = json.Unmarshal(data, &cocktailsList)
	if err != nil {
		fmt.Println("error:", err)
	}

	s.cocktailsList = cocktailsList
	s.ById = convertCocktailsListIntoById(cocktailsList)
}

func convertCocktailsListIntoById(cocktails []Cocktail) map[int]Cocktail {
	byId := make(map[int]Cocktail)
	for _, c := range cocktails {
		byId[c.Id] = c
	}
	return byId
}

func convertCocktailsListInByIngredients(cocktails []Cocktail) map[string][]int {

	return nil
}
