package handy

import (
	"fmt"
	"log"
	"net/http"
)

var (
	Config = map[string]interface{}{
		"TemplatePath": "templates",
		"CookieSecret": "handy_secret_cookie",
		"Address":      "",
		"Port":         "8080",
		"Debug":        false,
	}
	driverName, dataSourceName string
)

type Application struct {
	Router *Router
}

func (app *Application) New(config map[string]interface{}) *Application {
	if config != nil {
		Config = config
	}
	loadTemplate()
	application := &Application{Router: newRouter()}
	return application
}

func (app *Application) Connection(dsn, conn string) {
	driverName = dsn
	dataSourceName = conn
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
	listen := fmt.Sprintf("%s:%s", address, port)
	http.Handle("/", app.Router)
	log.Print("Listening on " + listen + "...")
	log.Fatal(http.ListenAndServe(listen, nil))
}
