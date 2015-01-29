package main

import _ "github.com/mattn/go-sqlite3"

import (
	"github.com/go-code/her"
	"github.com/go-code/her/samples/book/handler"
)

func main() {
	app := her.NewApplication()

	app.Database.Connection("sqlite", "sqlite3", "./book.s3db")

	app.Route.Handle("/", handler.Book.HomeHandler)
	app.Route.Handle("/create", handler.Book.CreateHandler)
	app.Route.Handle("/del/{id:[0-9]+}", handler.Book.DeleteHandler)
	app.Route.Handle("/static/{path:.*}", her.StaticFileHandler("static"))

	app.Start()
}
