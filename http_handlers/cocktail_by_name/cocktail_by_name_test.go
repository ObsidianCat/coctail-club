package cocktail_by_name

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

func TestCocktailByNameRoute(t *testing.T) {
	store.StoreInit()
	router := server.SetupRouter()

	t.Run("Respond with cocktail details", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/name/mojito", nil)
		router.ServeHTTP(w, req)
		var result store.Cocktail

		require.Equal(200, w.Code)
		err := json.NewDecoder(w.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Unable to parse response")
		}
		require.Equal(result.Name, "Mojito")
	})

	t.Run("Respond with error message when cocktail with given name does not exist", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/name/mokito", nil)
		router.ServeHTTP(w, req)

		require.Equal(404, w.Code)

		var result map[string]string
		err := json.NewDecoder(w.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Unable to parse response")
		}
		require.Equal(result[common.ErrorMessageKey], "Cocktail with this name not found")

	})
}
