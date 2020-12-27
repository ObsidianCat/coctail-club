package common

// ErrorHeaderName is used for error headers on response
const ErrorHeaderName = "X-Error-Description"

// ErrorMessageKey is used for error message in response body
const ErrorMessageKey = "error"

// CocktailDbURLSearchByName API URL
const CocktailDbURLSearchByName = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="

// CocktailDbURLLookupByID API URL
const CocktailDbURLLookupByID = "https://www.thecocktaildb.com/api/json/v1/1/lookup.php?i="

// CocktailDbURLSeacrhByIngredient  API URL
const CocktailDbURLSeacrhByIngredient = "https://www.thecocktaildb.com/api/json/v1/1/filter.php?i="

// CocktailCollectionFilePath default path for collection file
const CocktailCollectionFilePath = "./cocktails_collection.json"
