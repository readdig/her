package main

import (
	"github.com/go-web-framework/handy"
	"github.com/go-web-framework/handy/samples/book/handler"
	_ "github.com/mattn/go-sqlite3"
)

var (
	config = map[string]interface{}{
		"TemplatePath": "templates",
		"CookieSecret": "handy_secret_cookie",
		"Address":      "",
		"Port":         "8080",
		"Debug":        true,
	}
	application *handy.Application
)

func main() {
	app := application.New(config)

	app.Connection("sqlite3", "./book.s3db")

	app.Router.HandleFunc("/", handler.Book.HomeHandler)
	app.Router.HandleFunc("/static/{path:.*}", handy.StaticHandler)
	app.Router.HandleFunc("/create", handler.Book.CreateHandler)

	app.Start()
}
