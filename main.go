package main

import (
	"cocktail-club/server"
	"cocktail-club/store"
)

func main() {
	store.StoreInit()
	r := server.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (localhost:8080)
}
