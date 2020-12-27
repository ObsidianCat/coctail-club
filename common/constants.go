package common

// ErrorHeaderName is used for error headers on response
const ErrorHeaderName = "X-Error-Description"

// ErrorMessageKey is used for error message in response body
const ErrorMessageKey = "error"

//
const CocktailDbUrlSearchByName = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
const CocktailDbUrlLookupById = "https://www.thecocktaildb.com/api/json/v1/1/lookup.php?i="
const CocktailDbUrlSeacrhByIngredient = "https://www.thecocktaildb.com/api/json/v1/1/filter.php?i="
const CocktailCollectionFilePath = "./cocktails_collection.json"
