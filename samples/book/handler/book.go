package handler

import (
	"github.com/go-framework/web"
	"github.com/go-framework/web/samples/book/form"
	"github.com/go-framework/web/samples/book/model"
)

type bookHandler struct{}

var (
	Book      = &bookHandler{}
	bookModel = &model.Book{}
)

func (h *bookHandler) HelloHandler(val string) string {
	return "hello " + val
}

func (h *bookHandler) HomeHandler(ctx *web.Context) {
	books := bookModel.GetAll()
	tmpl := map[string]interface{}{}
	tmpl["books"] = books

	ctx.Render("index.html", tmpl)
}

func (h *bookHandler) CreateHandler(ctx *web.Context) {
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
