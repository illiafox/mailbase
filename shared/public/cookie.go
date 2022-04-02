package public

import (
	"errors"
	"time"
)

var Cookie = struct {
	CookieError error
	MaxAge      int64
}{
	CookieError: errors.New("cookie error! Please, try <a href='/login'>Login</a> again"),
	MaxAge:      int64(time.Hour * 24 * 7), // 7 days
}
