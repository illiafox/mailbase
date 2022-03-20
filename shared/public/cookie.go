package public

import "errors"

var Cookie = cookie{
	CookieError: errors.New("cookie error! Please, enable them - we store only credentials"),
}

type cookie struct {
	CookieError error
}
