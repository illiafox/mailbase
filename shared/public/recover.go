package public

import "errors"

var Forgot = forgot{
	SamePassword: errors.New("this is your actual password :)<br><a href='/login'>login</a>"),
}

type forgot struct {
	SamePassword error
}
