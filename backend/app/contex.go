package octopus

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"html/template"
	"io"
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
	Values         map[any]any
}

func (c *Context) BodyParser(out interface{}) error {
	body, err := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return err
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	return nil
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
