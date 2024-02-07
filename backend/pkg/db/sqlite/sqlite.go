package sqlite

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() *sql.DB {
	_, errorNoFile := os.Stat("./pkg/db/sqlite/social-network.db")

	DB, err := sql.Open("sqlite3", "./pkg/db/sqlite/social-network.db")
	if err != nil {
		log.Println(err)
	}
	
	if errorNoFile != nil {
		sqlCode, ERR := os.ReadFile("./pkg/db/sqlite/init.sql")
		if ERR != nil {
			log.Fatal(ERR)
		}
		_, erp := DB.Exec(string(sqlCode))
		if erp != nil {
			log.Fatal(erp)
		}
	}

	return DB
}
