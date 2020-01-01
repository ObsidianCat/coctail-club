package server

import (
	"cocktail-club/common"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCocktailByIngredientRoute(t *testing.T) {
	common.StoreInit()
	router := SetupRouter()

	t.Run("Respond with cocktails by ingredient", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/ingredient?sugar", nil)
		router.ServeHTTP(w, req)
		var results []common.Cocktail

		require.Equal(200, w.Code)
		err := json.NewDecoder(w.Body).Decode(&results)
		if err != nil {
			t.Fatalf("Unable to parse response")
		}
		require.Equal(len(results), 2)
		require.Equal(results[0].ID, 2)
		require.Equal(results[1].ID, 3)
	})

	t.Run("Respond with error message in body when no cocktails contain search quarry ingredient", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/ingredient?sugarrr", nil)
		router.ServeHTTP(w, req)

		require.Equal(404, w.Code)

		var results map[string]string
		err := json.NewDecoder(w.Body).Decode(&results)
		if err != nil {
			t.Fatalf("Unable to parse response")
		}
		require.Equal(results[common.ErrorMessageKey], "No results for this ingredient")

	})
}
