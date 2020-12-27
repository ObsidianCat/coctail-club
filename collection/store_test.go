package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreCreation(t *testing.T) {
	require := assert.New(t)
	StoreInit("../fixtures/cocktails_collection.json")
	storePointer = GetStore()
	require.NotNil(storePointer)
	require.Equal(len(storePointer.Cocktails), 3, "Should contain 3 cocktail")
	require.Equal(storePointer.Cocktails[0].Name, "Mojito", "First cocktail name should be equal Mojito")
}
