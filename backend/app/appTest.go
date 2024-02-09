package octopus

import (
	"backend/pkg/db/sqlite"
	"database/sql"
)

var migrate = sqlite.Migrations{}
var AppTest = New(migrate)

var Db = func() *sql.DB {
	conn, err := sql.Open("sqlite3", "../db/sqlite/app_database.sqlite")
	if err != nil {
		panic(err)
	}
	return conn
}()

var _ = func() error {
	// AppTest.Run(":8888")
	AppTest.UseDb(Db)

	return nil
}()