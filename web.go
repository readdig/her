package web

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
)

var (
	Config = map[string]interface{}{
		"TemplatePath": "templates",
		"CookieSecret": "web_secret_cookie",
		"Address":      "",
		"Port":         "8080",
		"Debug":        false,
	}
	driverName, dataSourceName string
)

type Application struct {
	Route *Router
}

func (app *Application) New(config map[string]interface{}) *Application {
	if config != nil {
		Config = config
	}
	application := &Application{Route: newRouter()}
	return application
}

func (app *Application) Connection(dsn, conn string) {
	driverName = dsn
	dataSourceName = conn
}

func (app *Application) FuncMap(tmplFunc map[string]interface{}) {
	if len(tmplFunc) > 0 {
		for k, v := range tmplFunc {
			funcMap[k] = v
		}
	}
}

func (app *Application) Start() {
	address, ok := Config["Address"].(string)
	if !ok {
		address = ""
	}
	port, ok := Config["Port"].(string)
	if !ok {
		port = "8080"
	}
	debug, ok := Config["Debug"].(bool)
	if !ok {
		debug = false
	}

	listen := fmt.Sprintf("%s:%s", address, port)
	mux := http.NewServeMux()
	if debug {
		mux.Handle("/debug/pprof", http.HandlerFunc(pprof.Index))
		mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		mux.Handle("/debug/pprof/block", pprof.Handler("block"))
		mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	}
	mux.Handle("/", app.Route)

	l, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Listening on " + listen + "...")
	log.Fatal(http.Serve(l, mux))
}
