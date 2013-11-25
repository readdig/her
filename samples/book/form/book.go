package form

import (
	"github.com/go-web-framework/handy"
)

type BookForm struct {
	Form     *handy.Form
	UserName *handy.TextField
	Content  *handy.TextAreaField
}

func NewBookForm(ctx *handy.Context) *BookForm {
	form := &BookForm{}
	form.UserName = handy.NewTextField("username", "用户名", "", handy.Required{}, handy.Length{Min: 3, Max: 10})
	form.Content = handy.NewTextAreaField("content", "内容", "", handy.Required{}, handy.Length{Min: 1, Max: 200})

	form.Form = handy.InitForm(ctx, form)
	return form
}
