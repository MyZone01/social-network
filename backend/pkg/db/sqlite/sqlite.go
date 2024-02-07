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
	currentDir, err := os.Getwd()
	fmt.Println(currentDir)

	// migrationDir := "/home/shaykh/Projects/Group_Proj/social-network/backend/pkg/db/migrations/"
	// databasePath := "/home/shaykh/Projects/Group_Proj/social-network/backend/pkg/db/sqlite/social-network.db"

	migrationDir := currentDir + "/pkg/db/migrations/"
	fmt.Println(migrationDir)
	databasePath := currentDir + "/pkg/db/sqlite/social-network.db"
	m, err := migrate.New("file://"+migrationDir, "sqlite://"+databasePath)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
	}
	defer m.Close()

	// Apply migrations (Up)
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	//     log.Fatal(err)
	// }

	// Rollback one migration (Down)
	if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
	}

	return DB
}
