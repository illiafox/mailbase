package cookie

import (
	"github.com/gorilla/sessions"
)

const _cookie_key = "2F423F4528482B4D6251655468576D5A7133743677397A24432646294A404E63A"

var (
	Store = sessions.NewCookieStore([]byte(_cookie_key))
)
