package templates

var (
	Error         = CompileTemplate("../shared/templates/error/index.html")
	Successful    = CompileTemplate("../shared/templates/succ/index.html")
	MailVerify    = CompileTemplate("../shared/templates/mailVerify/index.html")
	MailResetPass = CompileTemplate("../shared/templates/mailResetPass/index.html")

	Reset = CompileTemplate("../shared/templates/reset/index.html")
	Main  = CompileTemplate("../shared/templates/main/index.html")
)
