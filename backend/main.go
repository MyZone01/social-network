package main

import (
	octopus "backend/app"
	"backend/pkg/handlers"
)

func main() {
	//initialisation of the backend application
	app := octopus.New()
	// lunch all handlers
	handlers.HandleAll(app)
	if err := app.Run(":8081"); err != nil {
		panic(err)
	}

}
