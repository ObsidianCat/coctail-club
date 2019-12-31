package main

import (
	"cocktail-club/common"
	"cocktail-club/server"
)

func main() {
	common.StoreInit()
	r := server.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
