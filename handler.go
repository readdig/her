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

	isContext := false
	in := make([]reflect.Value, handlerType.NumIn())
	for i := 0; i < handlerType.NumIn(); i++ {
		argType := handlerType.In(i)

		if argType.Kind() == reflect.Ptr {
			if i == 0 && argType.Elem().Name() == "Context" {
				in[i] = reflect.ValueOf(ctx)
				isContext = true
				continue
			}
		}

		if isContext {
			in[i] = reflect.ValueOf(vars[1:][i-1])
		} else {
			in[i] = reflect.ValueOf(vars[1:][i])
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
			ctx.NotFound()
			return
		}
		if info.IsDir() {
			ctx.Forbidden()
			return
		}
		http.ServeFile(ctx.ResponseWriter, ctx.Request, path)
	}
}
