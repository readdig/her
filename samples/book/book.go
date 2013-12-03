package main

import (
	"github.com/go-framework/handy"
	"github.com/go-framework/handy/samples/book/handler"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
)

var (
	config = map[string]interface{}{
		"TemplatePath": "templates",
		"CookieSecret": "handy_secret_cookie",
		"Address":      "",
		"Port":         "8080",
		"Debug":        true,
	}
	fucnMap = map[string]interface{}{
		"hello": func(text string) template.HTML {
			return template.HTML(text)
		},
	}
	application *handy.Application
)

func main() {
	app := application.New(config)

	app.Connection("sqlite3", "./book.s3db")
	app.FuncMap(fucnMap)

	app.Router.HandleFunc("/", handler.Book.HomeHandler)
	app.Router.HandleFunc("/static/{path:.*}", handy.StaticHandler)
	app.Router.HandleFunc("/create", handler.Book.CreateHandler)

	app.Start()
}
