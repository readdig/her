package her

import (
	"fmt"
	"html/template"
	"strings"
)

type Field interface {
	Label(attrs ...string) template.HTML
	Render(attrs ...string) template.HTML
	Validate() bool
	Text() string
	Name() string
	Value() string
	SetValue(value string)
	HasErrors() bool
	AddError(err string)
	Errors() []string
	ValidatorMessage(attrs ...string) template.HTML
}

type BaseField struct {
	name       string
	text       string
	value      string
	errors     []string
	validators []Validator
}

func (field *BaseField) Label(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}
	return template.HTML(fmt.Sprintf(`<label for="%s"%s>%s</label>`, field.name, attrsStr, field.text))
}

func (field *BaseField) HasErrors() bool {
	return len(field.errors) > 0
}

func (field *BaseField) Render(attrs ...string) template.HTML {
	return template.HTML("")
}

func (field *BaseField) Text() string {
	return field.text
}

func (field *BaseField) Name() string {
	return field.name
}

func (field *BaseField) AddError(err string) {
	field.errors = append(field.errors, err)
}

func (field *BaseField) Errors() []string {
	return field.errors
}

func (field *BaseField) ValidatorMessage(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}

	result := ""
	for _, err := range field.errors {
		result += fmt.Sprintf(`<span%s>%s %s</span>`, attrsStr, field.text, err)
	}

	return template.HTML(result)
}

func (field *BaseField) Validate() bool {
	for _, validator := range field.validators {
		if _, ok := validator.(Required); ok {
			if ok, message := validator.CleanData(field.value); !ok {
				field.errors = append(field.errors, message)
				return false
			}
		}
	}

	result := true

	for _, validator := range field.validators {
		if ok, message := validator.CleanData(field.value); !ok {
			result = false
			field.errors = append(field.errors, message)
		}
	}

	return result
}

func (field *BaseField) Value() string {
	return field.value
}

func (field *BaseField) SetValue(value string) {
	field.value = value
}

// textarea
type TextAreaField struct {
	BaseField
}

func (field *TextAreaField) Render(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}

	return template.HTML(fmt.Sprintf(`<textarea id="%s" name="%s"%s>%s</textarea>`, field.name, field.name, attrsStr, field.value))
}

func NewTextAreaField(name string, text string, value string, validators ...Validator) *TextAreaField {
	field := TextAreaField{}
	field.name = name
	field.text = text
	field.value = value
	field.validators = validators

	return &field
}

// select
type Choice struct {
	Text  string
	Value string
}

type SelectField struct {
	BaseField
	Choices []Choice
}

func (field *SelectField) Render(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}

	options := ""
	for _, choice := range field.Choices {
		selected := ""
		if choice.Value == field.value {
			selected = " selected"
		}
		options += fmt.Sprintf(`<option value="%s"%s>%s</option>`, choice.Value, selected, choice.Text)
	}

	return template.HTML(fmt.Sprintf(`<select id="%s" name="%s"%s>%s</select>`, field.name, field.name, attrsStr, options))
}

func NewSelectField(name string, text string, choices []Choice, defaultValue string, validators ...Validator) *SelectField {
	field := SelectField{}
	field.name = name
	field.text = text
	field.value = defaultValue
	field.Choices = choices
	field.validators = validators

	return &field
}

// input hidden
type HiddenField struct {
	BaseField
}

func (field *HiddenField) Render() template.HTML {
	return template.HTML(fmt.Sprintf(`<input type="hidden" value="%s" name="%s" id="%s">`, field.value, field.name, field.name))
}

func NewHiddenField(name string, value string) *HiddenField {
	field := HiddenField{}
	field.name = name
	field.value = value

	return &field
}

// input text
type TextField struct {
	BaseField
}

func (field *TextField) Render(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}
	return template.HTML(fmt.Sprintf(`<input type="text" value="%s" name="%s" id="%s"%s>`, field.value, field.name, field.name, attrsStr))
}

func NewTextField(name string, text string, value string, validators ...Validator) *TextField {
	field := TextField{}
	field.name = name
	field.text = text
	field.value = value
	field.validators = validators

	return &field
}

// input  password
type PasswordField struct {
	BaseField
}

func (field *PasswordField) Render(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}
	return template.HTML(fmt.Sprintf(`<input type="password" name="%s" id="%s"%s>`, field.name, field.name, attrsStr))
}

func NewPasswordField(name string, text string, validators ...Validator) *PasswordField {
	field := PasswordField{}
	field.name = name
	field.text = text
	field.validators = validators

	return &field
}

// Radio
type RadioField struct {
	BaseField
}

func (field *RadioField) Render(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}
	return template.HTML(fmt.Sprintf(`<input type="radio" name="%s" id="%s" value="%s"%s>`, field.name, field.name, field.value, attrsStr))
}

func NewRadioField(name, text, value string, validators ...Validator) *RadioField {
	field := RadioField{}
	field.name = name
	field.text = text
	field.value = value
	field.validators = validators

	return &field
}

// checkboes
type CheckField struct {
	BaseField
}

func (field *CheckField) Render(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}
	return template.HTML(fmt.Sprintf(`<input type="checkbox" name="%s" id="%s" value="%s"%s>`, field.name, field.name, field.value, attrsStr))
}

func NewCheckField(name, text, value string, validators ...Validator) *CheckField {
	field := CheckField{}
	field.name = name
	field.text = text
	field.value = value
	field.validators = validators

	return &field
}

//SubmitField

type SubmitField struct {
	BaseField
}

func (field *SubmitField) Render(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}
	return template.HTML(fmt.Sprintf(`<input type="submit" name="%s" id="%s" value="%s"%s>`, field.name, field.name, attrsStr, field.value))
}

func NewSubmitField(name, text, value string, validators ...Validator) *SubmitField {

	field := SubmitField{}
	field.name = name
	field.text = text
	field.value = value
	field.validators = validators

	return &field
}

//FileField
type FileField struct {
	BaseField
}

func (field *FileField) Render(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		attrsStr = " " + strings.Join(attrs, " ")
	}
	return template.HTML(fmt.Sprintf(`<input type="file" name="%s" id="%s"%s>`, field.name, field.name, attrsStr))
}

func NewFileField(name, text string, validators ...Validator) *FileField {

	field := FileField{}
	field.name = name
	field.text = text
	field.validators = validators

	return &field
}
