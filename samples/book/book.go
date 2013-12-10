package main

import _ "github.com/mattn/go-sqlite3"

import (
	"github.com/go-web-framework/handy"
	"github.com/go-web-framework/handy/samples/book/handler"
)

var (
	config = map[string]interface{}{
		"TemplatePath": "templates",
		"Address":      "127.0.0.1",
		"Port":         "8080",
		"Debug":        true,
	}
	application = &handy.Application{}
)

func main() {
	app := application.New(config)

	app.Connection("sqlite3", "./book.s3db")

	app.Route.Handle("/", handler.Book.HomeHandler)
	app.Route.Handle("/static/{path:.*}", handy.StaticFileHandler("static"))
	app.Route.Handle("/create", handler.Book.CreateHandler)

	app.Start()
}
