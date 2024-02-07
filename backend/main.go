package main

import (
	octopus "backend/app"
	"backend/pkg/handlers"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//initialisation of the backend application
	app := octopus.New()

	fmt.Println(app)

	// lunch all handlers
	handlers.HandleAll(app)
	if err := app.Run(":8081"); err != nil {
		panic(err)
	}

}
