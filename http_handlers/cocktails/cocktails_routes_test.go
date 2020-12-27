package cocktails

import (
	"cocktail-club/collection"
	"cocktail-club/common"
	"cocktail-club/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCocktailsByIngredientRoute(t *testing.T) {
	collection.StoreInit()
	router := server.SetupRouter()

	t.Run("Respond with cocktails list", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/ingredient?sugar", nil)
		router.ServeHTTP(w, req)
		var result []common.Cocktail

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

func TestCocktailsByNameRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/cocktail/name/:name", GetByName)

	t.Run("Respond with cocktail details", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/name/mojito", nil)
		router.ServeHTTP(w, req)
		var result []common.Cocktail

		require.Equal(200, w.Code)
		err := json.NewDecoder(w.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Unable to parse response")
		}
		require.Equal(result[0].Name, "Mojito")
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

func TestCocktailByIdRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/cocktail/id/:name", GetById)

	t.Run("Respond with cocktail details", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/id/11000", nil)
		router.ServeHTTP(w, req)
		var result []common.Cocktail

		require.Equal(200, w.Code)
		err := json.NewDecoder(w.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Unable to parse response")
		}
		require.Equal(result[0].Name, "Mojito")
	})

	t.Run("Respond with error message when cocktail with given name does not exist", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/id/test123", nil)
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
