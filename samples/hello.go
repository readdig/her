package main

import (
	"github.com/go-web-framework/handy"
)

var (
	application = &handy.Application{}
)

func main() {
	app := application.New(nil)
	app.Route.Handle("/", func() string {
		return "hello world!"
	})
	app.Route.Handle("/hello/{val}", func(val string) string {
		return "hello " + val
	})
	app.Route.Handle("/hi/{val}", func(ctx *handy.Context, val string) {
		ctx.WriteString("hi " + val)
	})
	app.Start()
}
