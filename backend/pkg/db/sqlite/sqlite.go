package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	migrationDir := "/home/shaykh/Projects/Group_Proj/social-network/backend/pkg/db/migrations/"
	databasePath := "/home/shaykh/Projects/Group_Proj/social-network/backend/pkg/db/sqlite/social-network.db"
	m, err := migrate.New("file://"+migrationDir, "sqlite://"+databasePath)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	if err, erro := m.Close(); err != nil || erro != nil {
		log.Fatal(err, erro)
	}

	return DB
}
