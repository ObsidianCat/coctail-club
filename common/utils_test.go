package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransformApiBytesToCtails(t *testing.T) {
	require := require.New(t)

	data, _ := ReadDataFileWithPathFromCallerFile("../fixtures/api_cocktails_search_results.json")
	result := TransformAPIBytesToCtails(data)
	require.Equal(3, len(result), "should return lsit of 3 cocktails")
	require.Equal([3]string{"Mojito", "Mojito #3", "Blueberry Mojito"}, [3]string{result[0].Name, result[1].Name, result[2].Name},
		"Cocktail names should match")
}
