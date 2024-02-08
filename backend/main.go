package main

import (
	octopus "backend/app"
	"backend/pkg/db/sqlite"
	"backend/pkg/handlers"
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
		if arg == "-up" || arg == "-down" || arg[:3] == "-to" || arg == "-upall" || arg == "-downall" {
			migrate.Migration = true
			if len(strings.Split(arg, "=")) == 2 {
				version, err := strconv.Atoi(strings.Split(arg, "=")[1])
				if err != nil || version == 0 {
					log.Println("incorrect version")
				} else {
					migrate.Target = true
					migrate.Version = version
					migrate.Action = strings.Split(arg, "=")[0]
				}
			} else {
				migrate.Target = true
				migrate.Action = arg
			}
		} else {
			migrate.Migration = false
		}
		break
	}
	//initialisation of the backend application
	app := octopus.New(migrate)

	// lunch all handlers
	handlers.HandleAll(app)
	// app.GET("/", func(ctx *octopus.Context) {
	// 	newUser := models.User{
	// 		Email:       "ma@dca.co",
	// 		Password:    "cdecqce",
	// 		FirstName:   "cewc",
	// 		LastName:    "cewc",
	// 		DateOfBirth: time.Time{},
	// 		AvatarImage: "cewc",
	// 		Nickname:    "cewc",
	// 		AboutMe:     "cewc",
	// 		IsPublic:    true,
	// 	}
	// 	err := newUser.Create(ctx.Db.Conn)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// })

	if err := app.Run(":8081"); err != nil {
		panic(err)
	}

}
