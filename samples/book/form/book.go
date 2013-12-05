package form

import (
	"github.com/go-framework/web"
)

type BookForm struct {
	Form     *web.Form
	UserName *web.TextField
	Content  *web.TextAreaField
}

func NewBookForm(ctx *web.Context) *BookForm {
	form := &BookForm{}
	form.UserName = web.NewTextField("username", "用户名", "", web.Required{}, web.Length{Min: 3, Max: 10})
	form.Content = web.NewTextAreaField("content", "内容", "", web.Required{}, web.Length{Min: 1, Max: 200})

	form.Form = web.InitForm(ctx, form)
	return form
}
