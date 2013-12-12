package handy

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	funcMap   template.FuncMap
	templates *template.Template
)

func init() {
	funcMap = make(template.FuncMap)
}

func loadTemplate() *template.Template {
	templatePath := Config.Get("TemplatePath").String()
	if templatePath == "" {
		return nil
	}
	for k, v := range templateFuncMap() {
		funcMap[k] = v
	}
	t, err := buildTemplate(templatePath, funcMap)
	if err != nil {
		log.Printf("Can't read template file %v,", err)
	}
	return t
}

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"set": func(renderArgs map[string]interface{}, key string, value interface{}) template.HTML {
			renderArgs[key] = value
			return template.HTML("")
		},
		"append": func(renderArgs map[string]interface{}, key string, value interface{}) template.HTML {
			if renderArgs[key] == nil {
				renderArgs[key] = []interface{}{value}
			} else {
				renderArgs[key] = append(renderArgs[key].([]interface{}), value)
			}
			return template.HTML("")
		},
		// Replaces newlines with <br>
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
		// Skips sanitation on the parameter.  Do not use with dynamic data.
		"raw": func(text string) template.HTML {
			return template.HTML(text)
		},
		"datetime": func(date time.Time, format string) string {
			return date.Format(format)
		},
		"xsrf_form_html": genTokenHTML,
	}
}

func buildTemplate(dir string, funcMap template.FuncMap) (*template.Template, error) {
	var t *template.Template
	return t, filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			filetext, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			text := string(filetext)
			name := path[len(dir)+1:]
			name = strings.Replace(name, `\`, `/`, -1)

			var tmpl *template.Template
			if t == nil {
				t = template.New(name)
			}
			if name == t.Name() {
				tmpl = t
			} else {
				tmpl = t.New(name)
			}
			_, err = tmpl.Funcs(funcMap).Parse(text)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
