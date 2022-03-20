package templates

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

func CompileTemplate(filepath ...string) *ExecTemplate {
	t, err := template.New("index.html").Funcs(
		template.FuncMap{
			"nlbr": func(text string) template.HTML {
				return template.HTML(strings.ReplaceAll(template.HTMLEscapeString(text), "\n", "<br>"))
			},
		},
	).ParseFiles(filepath...)
	return &ExecTemplate{Tmpl: template.Must(t, err)}
}

type ExecTemplate struct {
	Tmpl *template.Template
}

func (ex *ExecTemplate) WriteAnyCode(w http.ResponseWriter, statusCode int, element ...any) error {
	w.WriteHeader(statusCode)
	return ex.WriteAny(w, element...)
}

func (ex *ExecTemplate) WriteAny(w http.ResponseWriter, element ...any) error {
	return ex.Tmpl.Execute(w, template.HTML(fmt.Sprint(element...)))
}

func (ex *ExecTemplate) WriteBytes(writer io.Writer, element ...any) error {

	return ex.Tmpl.Execute(writer, template.HTML(fmt.Sprint(element...)))

}
