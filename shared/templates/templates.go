package templates

var (
	Error      = CompileTemplate("../shared/templates/error/index.html")
	Successful = CompileTemplate("../shared/templates/succ/index.html")
	Reset      = CompileTemplate("../shared/templates/reset/index.html")
	Main       = CompileTemplate("../shared/templates/main/index.html")
)

var Mail = struct {
	Verify, ResetPass *ExecTemplate
}{
	Verify:    CompileTemplate("../shared/templates/mails/VerifyAddress/index.html"),
	ResetPass: CompileTemplate("../shared/templates/mails/ResetPass/index.html"),
}
