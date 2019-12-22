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
		var ingridients []string
		for _, ing := range c.Ingredients {
			ings := strings.Fields(ing)
			fmt.Println(ings, len(ings))
			ingridients = append(ingridients, ings...)
		}
	}
	return byIngredient
}
