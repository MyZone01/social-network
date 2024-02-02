package octopus

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type db struct {
	Conn *sql.DB
}

type Context struct {
	Db             *db
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	handlers       []HandlerFunc
	index          int
}

func (c *Context) BodyParser(out interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(&out)
}

func (c *Context) JSON(data interface{}) error {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(c.ResponseWriter).Encode(data)
}

func (c *Context) Next() {
	if c.index < len(c.handlers) {
		handler := c.handlers[c.index]
		c.index++
		handler(c)
	}
}

func (c *Context) Render(path string, data interface{}) error {
	tp, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	return tp.Execute(c.ResponseWriter, data)
}

func (c *Context) Status(code int) *Context {
	c.ResponseWriter.WriteHeader(code)
	return c
}

func (c *Context) WriteString(s string) (int, error) {
	return c.ResponseWriter.Write([]byte(s))
}

func (db *db) InsertData(tableName string, columns []string, values ...interface{}) error {
	if db.Conn == nil {
		return fmt.Errorf("the app have no database")
	}
	if len(columns) != len(values) {
		return fmt.Errorf("le nombre de colonnes ne correspond pas au nombre de valeurs")
	}
	query := fmt.Sprintf("INSERT INTO %s (%s", tableName, columns[0])
	for i := 1; i < len(columns); i++ {
		query += ", " + columns[i]
	}
	query += ")"
	placeholders := make([]string, len(columns))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	query += " VALUES (" + placeholders[0]
	for i := 1; i < len(placeholders); i++ {
		query += ", " + placeholders[i]
	}
	query += ")"
	stmt, err := db.Conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}
	return nil
}
