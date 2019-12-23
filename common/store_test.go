package common

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var cocktailsListMock = []Cocktail{
	Cocktail{
		Id:          1,
		Name:        "Espresso Martini",
		Ingredients: []string{"Vodka", "Kahlua", "Espresso Coffee", "Sugar Syrup"},
		Preparation: string("Combine all ingredients in a shaker filled with ice & shake well.\nStrain the mixture into a chilled cocktail Martini glass.\nCarefully place three coffee beans on top of the drink for garnish & serve."),
		Url:         string("https://thebarcabinet.com/recipes/vodka-cocktails/espresso-martini/"),
	},
	Cocktail{
		Id:          2,
		Name:        "Mojito",
		Ingredients: []string{"White Rum", "Sugar Syrup", "Lime Wedges", "Fresh Mint", "Soda Water"},
		Preparation: string("Place the Mint, Sugar Syrup & Lime wedges into a highball glass & lightly muddle the ingredients together. The Lime wedges & Mint leaves should be bruised to release their juices & essential oils.\nFill the glass with crushed ice, pour over the White Rum & stir.\nTop up with Soda Water & stir well from the bottom up.\nGarnish with a sprig of Mint & serve with a straw."),
		Url:         string("https://thebarcabinet.com/recipes/rum-cocktails/mojito/"),
	}}

func TestStoreCreation(t *testing.T) {
	require := assert.New(t)
	storePointer = GetStore()
	require.NotNil(storePointer)
	require.Equal(storePointer.ByName["mojito"], 2, "By name map return cocktail id")
}

func TestLoadingDataIntoStore(t *testing.T) {
	t.Run("Create ByIngredients structure", func(t *testing.T) {
		require := assert.New(t)

		result := convertCocktailsListInByIngredient(cocktailsListMock)
		require.NotNil(result)

		termResults1 := result["rum"]
		termResults2 := result["sugar"]
		require.NotNil(termResults1)
		require.NotNil(termResults2)
		require.Equal(len(termResults1), 1, "This ingredient used once")
		require.Equal(len(termResults2), 2, "This ingredient used twice")
		require.Equal(termResults1, []int{2})
		require.Equal(termResults2, []int{1, 2})

	})

	t.Run("Create ById structure", func(t *testing.T) {
		result := convertCocktailsListIntoById(cocktailsListMock)
		require.NotNil(t, result)
	})
}
