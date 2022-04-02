package templates

import (
	"html/template"
	"strings"
)

var (
	Error      = CompileTemplate("../shared/templates/common/error/index.html")
	Successful = CompileTemplate("../shared/templates/common/succ/index.html")

	Reset = CompileTemplate("../shared/templates/site/reset/index.html")
	Main  = CompileTemplate("../shared/templates/site/main/index.html")
)

var Admin = struct {
	Panel  *ExecTemplate
	Report *ExecTemplate
}{
	Panel: CompileTemplate("../shared/templates/admin/panel/index.html"),

	Report: CompileTemplateFuncs(template.FuncMap{
		"nlbr": func(text string) template.HTML {
			return template.HTML(strings.ReplaceAll(text, "\n", "<br>"))
		}}, "../shared/templates/admin/report/index.html"),
}

var Mail = struct {
	Verify, ResetPass, Answer *ExecTemplate
}{
	Verify:    CompileTemplate("../shared/templates/mails/VerifyAddress/index.html"),
	ResetPass: CompileTemplate("../shared/templates/mails/ResetPass/index.html"),
	Answer:    CompileTemplate("../shared/templates/mails/Answer/index.html"),
}
