package handler

import (
	"github.com/go-web-framework/handy"
	"github.com/go-web-framework/handy/samples/book/form"
	"github.com/go-web-framework/handy/samples/book/model"
)

type bookHandler struct{}

var (
	Book      = &bookHandler{}
	bookModel = &model.Book{}
)

func (h *bookHandler) HomeHandler(ctx *handy.Context) {
	books := bookModel.GetAll()
	tmpl := map[string]interface{}{}
	tmpl["books"] = books

	ctx.Render("index.html", tmpl)
}

func (h *bookHandler) CreateHandler(ctx *handy.Context) {
	form := form.NewBookForm(ctx)
	tmpl := map[string]interface{}{}

	if ctx.Request.Method == "POST" {
		if form.Form.Validate() {
			bookModel := &model.Book{}
			bookModel.UserName = form.UserName.Value()
			bookModel.Content = form.Content.Value()
			result := bookModel.Insert()
			if result {
				tmpl["success"] = "发布成功"
			}
		}
	}
	tmpl["form"] = form
	ctx.Render("create.html", tmpl)
}
