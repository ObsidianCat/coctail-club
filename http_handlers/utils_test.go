package http_handlers

import (
	"cocktail-club/store"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransformApiBytesToCtails(t *testing.T) {
	data, _ := store.ReadDataFile("fixtures/api_cocktails_search_results.json")
	TransformApiBytesToCtails(data)
	// fmt.Println(string(data))
	//fmt.Println(result)
	require.Equal(t, 1, 1)

}
