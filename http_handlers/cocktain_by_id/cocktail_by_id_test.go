package cocktain_by_id

import (
	"cocktail-club/common"
	"cocktail-club/store"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCocktailByIdRoute(t *testing.T) {
	router := gin.Default()
	router.GET("/cocktail/id/:name", CocktailById)

	t.Run("Respond with cocktail details", func(t *testing.T) {
		require := require.New(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/id/11000", nil)
		router.ServeHTTP(w, req)
		var result []store.Cocktail

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
