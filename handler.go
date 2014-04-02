package her

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

type Handler interface{}

func routeHandler(ctx *Context, handler Handler, vars []string) {
	handlerType := reflect.TypeOf(handler)
	if handlerType.Kind() != reflect.Func {
		panic("handler must be a callable func")
	}

	var in []reflect.Value
	if handlerType.NumIn() > 0 {
		argType := handlerType.In(0)
		if argType.Kind() == reflect.Ptr {
			if argType.Elem().Name() == "Context" {
				in = append(in, reflect.ValueOf(ctx))
			}
		}
		for _, arg := range vars[1:] {
			in = append(in, reflect.ValueOf(arg))
		}
	}
	ret := reflect.ValueOf(handler).Call(in)
	if len(ret) == 0 {
		return
	}
	sval := ret[0]
	var content []byte
	if sval.Kind() == reflect.String {
		content = []byte(sval.String())
	} else if sval.Kind() == reflect.Slice && sval.Type().Elem().Kind() == reflect.Uint8 {
		content = sval.Interface().([]byte)
	}
	ctx.SetHeader("Content-Length", strconv.Itoa(len(content)))
	ctx.Write(content)
}

func redirectHandler(url string, code int) Handler {
	return func(ctx *Context) {
		http.Redirect(ctx.ResponseWriter, ctx.Request, url, code)
	}
}

func StaticFileHandler(dir string) Handler {
	return func(ctx *Context, path string) {
		path = fmt.Sprintf("%s/%s", dir, path)
		info, err := os.Stat(path)
		if err != nil {
			ctx.Abort(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		if info.IsDir() {
			ctx.Abort(http.StatusForbidden, http.StatusText(http.StatusForbidden))
			return
		}
		http.ServeFile(ctx.ResponseWriter, ctx.Request, path)
	}
}
