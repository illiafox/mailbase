package templates

var (
	Error      = CompileTemplate("../shared/templates/error/index.html")
	Successful = CompileTemplate("../shared/templates/succ/index.html")
	Mail       = CompileTemplate("../shared/templates/mail/index.html")
	Main       = CompileTemplate("../shared/templates/main/index.html")
)
