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

### 安装
请确保Go环境已经安装，如未安装请参考 [Go 环境安装](http://golang.org/doc/install.html)，请安装最新版。

``` go
go get github.com/go-code/her
```

### 使用
```go
package main

import (
    "github.com/go-code/her"
)

var (
    application = &her.Application{}
)

func main() {
    app := application.New(nil)
    app.Route.Handle("/", func() string {
        return "hello world!"
    })
    app.Route.Handle("/hello/{val}", func(val string) string {
        return "hello " + val
    })
    app.Route.Handle("/hi/{val}", func(ctx *her.Context, val string) {
        ctx.WriteString("hi " + val)
    })
    app.Start()
}
```
启动程序访问8080端口，默认端口为8080

### 参考、使用项目
- gorilla [mux](https://github.com/gorilla/mux) 路由
- jimmykuu [wtforms](https://github.com/jimmykuu/wtforms) 表单
- fsnotify [fsnotify](https://github.com/robfig/fsnotify) 模版刷新

### 开发成员
John, Monkey
