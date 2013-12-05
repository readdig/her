package main

import (
	"github.com/go-framework/web"
	"github.com/go-framework/web/samples/book/handler"
	_ "github.com/mattn/go-sqlite3"
)

var (
	config = map[string]interface{}{
		"TemplatePath": "templates",
		"CookieSecret": "web_secret_cookie",
		"Address":      "127.0.0.1",
		"Port":         "8080",
		"Debug":        true,
	}
	application *web.Application
)

func main() {
	app := application.New(config)

	app.Connection("sqlite3", "./book.s3db")

	app.Route.Handle("/", handler.Book.HomeHandler)
	app.Route.Handle("/hello/{val}", handler.Book.HelloHandler)
	app.Route.Handle("/static/{path:.*}", web.StaticFileHandler("static"))
	app.Route.Handle("/create", handler.Book.CreateHandler)

	app.Start()
}
