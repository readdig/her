package main

import _ "github.com/mattn/go-sqlite3"

import (
	"github.com/go-code/her"
	"github.com/go-code/her/samples/book/handler"
)

var (
	config = map[string]interface{}{
		"TemplatePath": "templates",
		"Address":      "0.0.0.0",
		"XSRFCookies":  true,
		"CookieSecret": "book_secert",
		"Port":         8080,
		"Debug":        true,
	}
	application = &her.Application{}
)

func main() {
	app := application.New(config)

	app.Connection("sqlite3", "./book.s3db")

	app.Route.Handle("/", handler.Book.HomeHandler)
	app.Route.Handle("/static/{path:.*}", her.StaticFileHandler("static"))
	app.Route.Handle("/create", handler.Book.CreateHandler)

	app.Start()
}
