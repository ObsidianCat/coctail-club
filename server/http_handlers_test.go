package server

import (
	"cocktail-club/common"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCocktailByIngredientRoute(t *testing.T) {
	common.StoreInit()
	router := SetupRouter()

	t.Run("Create ByIngredients structure", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cocktail/ingredient?sugar", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		fmt.Println(w.Body.String())
		//assert.Equal(t, "pong", w.Body.String())
	})
}
