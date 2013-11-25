package handy

import (
	"net/http"
	"os"
)

type Handler interface {
	HandleRequest(*Context)
}

type HandlerFunc func(*Context)

func (f HandlerFunc) HandleRequest(c *Context) {
	f(c)
}

func redirectHandler(url string, code int) HandlerFunc {
	return func(ctx *Context) {
		http.Redirect(ctx.ResponseWriter, ctx.Request, url, code)
		return
	}
}

func StaticHandler(ctx *Context) {
	path := ctx.Request.URL.Path[len("/"):]
	info, err := os.Stat(path)
	if err != nil {
		http.Error(ctx.ResponseWriter, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if info.IsDir() {
		http.Error(ctx.ResponseWriter, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	http.ServeFile(ctx.ResponseWriter, ctx.Request, path)
}
