package public

import "errors"

var Login = struct {
	MailExist         error
	MailNotFound      error
	IncorrectPassword error
}{
	MailExist:         errors.New("email Already Exists"),
	MailNotFound:      errors.New("mail Not Found"),
	IncorrectPassword: errors.New("incorrect password"),
}
