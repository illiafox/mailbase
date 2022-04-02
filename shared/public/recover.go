package public

import "errors"

var Forgot = struct {
	SamePassword error
}{
	SamePassword: errors.New("this is your actual password :)<br><a href='/login'>login</a>"),
}
