package cookie

import (
	"github.com/gorilla/sessions"
	"net/http"
)

const cookie = "2F423F4528482B4D6251655468576D5A7133743677397A24432646294A404E63A"

var (
	Store *sessions.CookieStore
)

func init() {
	Store = sessions.NewCookieStore([]byte(cookie))
	Store.Options.Secure = true
	Store.Options.SameSite = http.SameSiteStrictMode
	Store.Options.MaxAge = 60 * 60 * 24 * 7 // 7 days
}
