package main

import (
	"github.com/go-code/her"
	"strconv"
)

type Info struct {
	Name        string
	Description string
}

func main() {
	app := her.NewApplication()
	app.Route.Handle("/", func() string {
		return "hello world!"
	})
	app.Route.Handle("/hello/{val}", func(val string) string {
		return "hello " + val
	})
	app.Route.Handle("/hi/{val}", func(ctx *her.Context, val int) {
		ctx.WriteString("hi " + strconv.Itoa(val))
	})

	app.Route.Handle("/par/{val}", func(ctx *her.Context) {
		ctx.WriteString("par: " + ctx.Params["val"])
	})

	app.Route.Handle("/api/her.xml", func(ctx *her.Context) {
		info := &Info{Name: "her", Description: "a web framework for golang"}
		ctx.Xml(info)
	})

	app.Route.Handle("/api/her.json", func(ctx *her.Context) {
		info := &Info{Name: "her", Description: "a web framework for golang"}
		ctx.Json(info)
	})

	app.Start()
}
