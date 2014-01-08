package her

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"
)

type Form struct {
	fields  map[string]Field
	context *Context
}

func InitForm(ctx *Context, f interface{}) *Form {
	form := Form{}
	form.context = ctx
	form.fields = make(map[string]Field)

	formv := reflect.ValueOf(f).Elem()

	for i := 0; i < formv.NumField(); i++ {
		field := formv.Field(i)

		f, ok := field.Interface().(Field)
		if ok {
			form.fields[f.Name()] = f
		}
	}
	return &form
}

func (form *Form) Validate() bool {
	result := true
	for name, field := range form.fields {
		field.SetValue(strings.TrimSpace(form.context.Request.FormValue(name)))
		if !field.Validate() {
			result = false
		}
	}
	return result
}

func (form *Form) Errors() []string {
	var errors []string
	for _, field := range form.fields {
		for _, err := range field.Errors() {
			errors = append(errors, err)
		}
	}
	return errors
}

func (form *Form) ValidatorSummary() template.HTML {
	result := ""
	for _, field := range form.fields {
		for _, err := range field.Errors() {
			result += fmt.Sprintf(`<li>%s %s</li>`, field.Text(), err)
		}
	}

	if result != "" {
		ul := fmt.Sprintf(`<ul class="validator-summary">%s</ul>`, result)
		return template.HTML(ul)
	}

	return template.HTML("")
}

func (form *Form) AddError(name, err string) {
	field := form.fields[name]
	field.AddError(err)
}

func (form *Form) Fields() map[string]Field {
	return form.fields
}
