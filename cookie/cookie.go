package cookie

import (
	"github.com/gorilla/sessions"
	"net/http"
)

// Hash sum
const cookie = "4DB2E134E453D713A48316A422DBDB812F2C79C2815F152C0147A3CB86864D53"

var (
	//	cookieByte = []byte(cookie)
	Store *sessions.CookieStore
)

func init() {
	Store = sessions.NewCookieStore([]byte(cookie))
	Store.Options.Secure = true
	Store.Options.SameSite = http.SameSiteStrictMode
	Store.Options.MaxAge = 60 * 60 * 24 * 7 // 7 days
}
