package common

//Cocktail item is a main data unit for the collection
type Cocktail struct {
	Name        string
	Ingredients []string
	ID          int
	Preparation string
	Image       string
}

type ApiCocktail struct {
	Name           string `json:"strDrink"`
	ID             string `json:"idDrink"`
	StrIngredient1 string `json:"strIngredient1"`
	StrIngredient2 string `json:"strIngredient2"`
	StrIngredient3 string `json:"strIngredient3"`
	StrIngredient4 string `json:"strIngredient4"`
	StrIngredient5 string `json:"strIngredient5"`
	strMeasure1    string `json:"strMeasure1"`
	strMeasure2    string `json:"strMeasure2"`
	StrMeasure3    string `json:"strMeasure3"`
	StrMeasure4    string `json:"strMeasure4"`
	StrMeasure5    string `json:"strMeasure5"`
	Preparations   string `json:"strInstructions"`
	Image          string `json:"strImageSource"`
}

type CocktailPreview struct {
	Name string `json:"strDrink"`
	ID   string `json:"idDrink"`
}
