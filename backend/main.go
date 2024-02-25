package main

import (
	octopus "backend/app"
	"backend/app/middleware/cors"
	"backend/pkg/config"
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
		if arg == "-up" || arg == "-down" || arg[:3] == "-to" || arg == "-up--all" || arg == "-down--all" {
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
	}

	//initialisation of the backend application
	app := octopus.New(migrate)
	database := sqlite.OpenDB(migrate)
	app.UseDb(database)

	// lunch all handlers
	app.Use(func(c *octopus.Context) {
		log.Println("\t\t[" + c.Request.URL.Path + c.Request.Method + "]")
		c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowCredentials: true,
		ExposedHeaders:   []string{},
		MaxAge:           86400,
	}))
	handlers.HandleAll(app)
	config.Sess.UseDB(app.Db.Conn)

	if err := app.Run(":8081"); err != nil {
		panic(err)
	}

}
