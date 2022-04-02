package templates

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func CompileTemplate(filepath ...string) *ExecTemplate {
	t, err := template.New("index.html").ParseFiles(filepath...)
	if err != nil {
		log.Fatalln(err)
	}
	return &ExecTemplate{Tmpl: t}
}

func CompileTemplateFuncs(m template.FuncMap, filepath ...string) *ExecTemplate {
	t, err := template.New("index.html").Funcs(m).ParseFiles(filepath...)
	if err != nil {
		log.Fatalln(err)
	}
	return &ExecTemplate{Tmpl: t}
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
