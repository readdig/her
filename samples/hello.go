package main

import (
	"github.com/go-code/her"
)

func main() {
	app := her.NewApplication()
	app.Route.Handle("/", func() string {
		return "hello world!"
	})
	app.Route.Handle("/hello/{val}", func(val string) string {
		return "hello " + val
	})
	app.Route.Handle("/hi/{val}", func(ctx *her.Context, val string) {
		ctx.WriteString("hi " + val)
	})

	app.Route.Handle("/par/{val}", func(ctx *her.Context) {
		ctx.WriteString("par: " + ctx.Params["val"])
	})
	app.Start()
}
