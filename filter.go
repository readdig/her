package her

import (
	"reflect"
)

type Filter interface{}

type ActionExecutingFilter struct {
	Context *Context
	Route   *Route
}

type ActionExecutedFilter struct {
	Context *Context
	Route   *Route
}

type ResultExecutingFilter struct {
	Context *Context
	Route   *Route
}

type ResultExecutedFilter struct {
	Context *Context
	Route   *Route
}

func routeHandlerFilter(ctx *Context, filter Filter) {
	if filter != nil {
		filterType := reflect.TypeOf(filter)
		if filterType.Kind() != reflect.Func {
			panic("filter must be a callable func")
		}
	}

}
