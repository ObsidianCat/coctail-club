package collection

import (
	"cocktail-club/common"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// Store contains all data for the service
type Store struct {
	dataPath  string
	Cocktails []common.Cocktail
}

var storePointer *Store

// StoreInit create new collection instance in memory and return reference to the newly created collection
func StoreInit(params ...string) *Store {
	path := common.CocktailCollectionFilePath
	if len(params) > 0 {
		path = params[0]
	}
	storePointer = &Store{dataPath: path}
	storePointer.LoadCocktails()
	return storePointer
}

// GetStore return reference to existing collection or create new collection and return its reference
func GetStore() *Store {
	if storePointer == nil {
		StoreInit()
	}
	return storePointer
}

func readDataFileWithPathFromRoot(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
		return nil, errors.New("cannot read file")
	}

	return data, nil
}

// LoadCocktails loads cocktail recipes from file into collection
func (s *Store) LoadCocktails() {
	// read file
	data, _ := readDataFileWithPathFromRoot(s.dataPath)

	// json data
	cocktailsList := []common.Cocktail{}

	err := json.Unmarshal(data, &cocktailsList)
	if err != nil {
		fmt.Println("error:", err)
	}

	//(*s) converting pointer of the receiver to the value of the receiver
	(*s).Cocktails = cocktailsList
}
