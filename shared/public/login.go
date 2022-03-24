package public

import "errors"

var Login = login{
	MailExist:         errors.New("email Already Exists"),
	MailNotFound:      errors.New("mail Not Found"),
	IncorrectPassword: errors.New("incorrect password"),
}

type login struct {
	MailExist         error
	MailNotFound      error
	IncorrectPassword error
}
