package main

import (
	octopus "backend/app"
	"backend/pkg/db/sqlite"
	"backend/pkg/handlers"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	args := os.Args[1:]

	migrate := sqlite.Migrations{}
	for _, arg := range args {
		if arg == "-up" || arg == "-down" || arg[:3] == "-to" /* || arg == "-upALL" || arg == "-downALL" */ {
			migrate.Migration = true
			if len(strings.Split(arg, "=")) == 2 {
				version, err := strconv.Atoi(strings.Split(arg, "=")[1])
				if err != nil || version == 0 {
					log.Println("incorrect version")
				} else {
					migrate.Target = true
					migrate.Version = version
				}
			}
		} else {
			migrate.Migration = false
		}
		break
	}
	//initialisation of the backend application
	app := octopus.New(migrate)

	fmt.Println(app)

	// lunch all handlers
	handlers.HandleAll(app)
	if err := app.Run(":8081"); err != nil {
		panic(err)
	}

}
