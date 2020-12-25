package cocktail_by_ingredient

import (
	"cocktail-club/common"
	"cocktail-club/server"
	"cocktail-club/store"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCocktailByIngredientRoute(t *testing.T) {
	store.StoreInit()
	router := server.SetupRouter()

	t.Run("Respond with cocktails list", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/ingredient?sugar", nil)
		router.ServeHTTP(w, req)
		var result []store.Cocktail

		require.Equal(200, w.Code)
		err := json.NewDecoder(w.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Unable to parse response")
		}
		require.Equal(len(result), 2)
		require.Equal(result[0].ID, 2)
		require.Equal(result[1].ID, 3)
	})

	t.Run("Respond with error message when no cocktails contain search quarry ingredient", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/ingredient?sugarrr", nil)
		router.ServeHTTP(w, req)

		require.Equal(404, w.Code)

		var result map[string]string
		err := json.NewDecoder(w.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Unable to parse response")
		}
		require.Equal(result[common.ErrorMessageKey], "No results for this ingredient")

	})
}
