package common

import "testing"
import "github.com/stretchr/testify/assert"

var cocktailsListMock = []Cocktail{
	Cocktail{
		Id:          1,
		Name:        "Espresso Martini",
		Ingredients: []string{"Vodka", "Tia Maria", "coffee", "sugar"},
	},
	Cocktail{
		Id:          2,
		Name:        "Clover Club",
		Ingredients: []string{"Jindea Gin", "Martini Bianco", "Chambord", "egg white", "sugar", "lemon"},
	}}

func TestStoreCreation(t *testing.T) {
	store := NewStorage()
	assert.NotNil(t, store)
}

func TestLoadingDataIntoStore(t *testing.T) {
	store := NewStorage()
	store.LoadCocktails()
	t.Run("Create ByIngredients structure", func(t *testing.T) {

	})

	t.Run("Create ById structure", func(t *testing.T) {
		result := convertCocktailsListIntoById(cocktailsListMock)
		assert.NotNil(t, result)
	})
}
