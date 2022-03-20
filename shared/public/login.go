package public

import "errors"

var Login = login{
	MailExist:         errors.New("email Already Exists"),
	LoginNotFound:     errors.New("login Not Found"),
	IncorrectPassword: errors.New("incorrect password"), // TODO: Add recover password feature
}

type login struct {
	MailExist         error
	LoginNotFound     error
	IncorrectPassword error
}
