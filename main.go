package main

import (
	"cocktail-club/collection"
	"cocktail-club/server"
)

func main() {
	collection.StoreInit()
	r := server.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (localhost:8080)
}
