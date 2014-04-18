her
=====
a web framework for golang

### 介绍
用 Go 实现的一个简单的 MVC 模式框架，目前支持：

* 路由/RESTFUL(route)
* 控制器(handler)
* 视图(templates)
* 表单(form)
* 静态文件(static)
* 数据库(database)

### 安装
请确保Go环境已经安装，如未安装请参考 [Go 环境安装](http://golang.org/doc/install.html)，请安装最新版。

``` go
go get github.com/go-code/her
```

### 快速使用
```go
package main

import (
    "github.com/go-code/her"
)

func main() {
    app := her.NewApplication()
    app.Route.Handle("/", func() string {
        return "hello world!"
    })
    app.Start()
}
```
启动程序访问8080端口，默认端口为8080

### her application

```go
app := her.NewApplication()
// app.Route and app.Database and app.Template
app.Start()
```

#### Config
```go
app := her.NewApplication(map[string]interface{}{
        "TemplatePath": "templates",
        "Address":      "0.0.0.0",
        "XSRFCookies":  true,
        "CookieSecret": "book_secert",
        "Port":         8080,
        "Debug":        true,
    })
```

#### Route ([mux](http://www.gorillatoolkit.org/pkg/mux))
```go
app.Route.Handle("/", func() {
// handle code
}) // get
app.Route.Handle("/", func() {
// handle code
}).Methods("POST") //post

app.Route.Handle("/{key}", func(key string) {
// handle code
})
app.Route.Handle("/{key}/{id:[0-9]+}", func(key string, id string) {
// handle code
})
```

#### Handler
```go
app.Route.Handle("/", func() string {
    return "hello world!"
})

app.Route.Handle("/", func(ctx *her.Context) {
    ctx.WriteString("hello world!")
})

app.Route.Handle("/{val}", func(ctx *her.Context, val string) {
    ctx.WriteString("hello world!")
})

app.Route.Handle("/par/{val}", func(ctx *her.Context) {
    ctx.WriteString("par: " + ctx.Params["val"])
})
```

#### Static file handler
```go
app.Route.Handle("/static/{path:.*}", her.StaticFileHandler("static")) // static 为静态文件目录
```

#### Database
```go
// import go-sqlite3
import _ "github.com/mattn/go-sqlite3"
// config
app.Database.Connection("sqlite", "sqlite3", "./book.s3db") // key, driver, data source
// use
DB = her.NewDB()
db := DB.Open("sqlite")
defer db.Close()
// sql code
```

#### Template FunMap
```go
app.Template.FuncMap(map[string]interface{}{
    "text": func(text string) template.HTML {
        return template.HTML(text)
    },
})
// use
{{text "her"}}
```

### 参考、使用项目
- gorilla [mux](https://github.com/gorilla/mux) 路由
- jimmykuu [wtforms](https://github.com/jimmykuu/wtforms) 表单
- fsnotify [fsnotify](https://github.com/robfig/fsnotify) 模版刷新

### 开发成员
John, Monkey

### LICENSE
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)