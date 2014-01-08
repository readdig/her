package form

import (
	"github.com/go-code/her"
)

type BookForm struct {
	Form     *her.Form
	UserName *her.TextField
	Content  *her.TextAreaField
}

func NewBookForm(ctx *her.Context) *BookForm {
	form := &BookForm{}
	form.UserName = her.NewTextField("username", "用户名", "", her.Required{}, her.Length{Min: 3, Max: 10})
	form.Content = her.NewTextAreaField("content", "内容", "", her.Required{}, her.Length{Min: 1, Max: 200})

	form.Form = her.InitForm(ctx, form)
	return form
}
