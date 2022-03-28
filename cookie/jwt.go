package cookie

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/illiafox/mailbase/shared/public"
	"net/http"
	"time"
)

const jwtKey = "8Zz5tw0Ionm3XPZZfN0NOml3z9FMfmpgXwovR9fp6ryDIoGRM8EPHAB6iHsc0fb"

var jwtKeyBytes = []byte(jwtKey)

// //

type SessionClaim struct {
	Key string
	jwt.StandardClaims
}

var Session session

type session struct {
}

func (session) SetClaim(w http.ResponseWriter, r *http.Request, key string) (string, error) {
	store, err := Store.Get(r, "credentials")
	if err != nil {
		//nolint:errorlint
		multi, ok := err.(securecookie.MultiError)
		if !ok || !multi.IsDecode() {
			return "", public.Cookie.CookieError
		}
	}

	claims := SessionClaim{
		Key: key,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: public.Cookie.MaxAge,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key, err = token.SignedString(jwtKeyBytes)
	if err != nil {
		return "", public.NewInternalWithError(err)
	}

	store.Values["key"] = key
	if store.Save(r, w) != nil {
		return "", public.Cookie.CookieError
	}
	return key, nil
}

func (session) GetClaim(r *http.Request) (string, error) {
	store, err := Store.Get(r, "credentials")
	if err != nil {
		return "", public.Cookie.CookieError
	}

	if store.IsNew {
		return "", public.Session.NoSession
	}

	jwtFromHeader, ok := store.Values["key"].(string)
	if !ok {
		return "", public.Session.NoSession
	}

	var claim SessionClaim
	token, err := jwt.ParseWithClaims(
		jwtFromHeader,
		&claim,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKeyBytes, nil
		},
	)

	if err != nil {
		return "", public.Session.NoSession
	}

	claims, ok := token.Claims.(*SessionClaim)
	if !ok {
		return "", public.Session.NoSession
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return "", public.Session.OldSession
	}

	return claims.Key, nil
}

func (session) DeleteClaim(w http.ResponseWriter, r *http.Request) error {
	store, err := Store.Get(r, "credentials")
	if err != nil {
		return public.Cookie.CookieError
	}
	store.Values = map[interface{}]interface{}{}

	if store.Save(r, w) != nil {
		return public.Cookie.CookieError
	}
	return nil
}
