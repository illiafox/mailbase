package cookie

import (
	"mailbase/shared/public"
	"net/http"
)

func SetSessionKey(w http.ResponseWriter, r *http.Request, key string) error {
	store, err := Store.Get(r, "credentials")
	if err != nil {
		return public.Cookie.CookieError
	}
	store.Values["key"] = key
	if store.Save(r, w) != nil {
		return public.Cookie.CookieError
	}
	return nil
}

func GetSessionKey(r *http.Request) (string, error) {
	store, err := Store.Get(r, "credentials")
	if err != nil {
		return "", public.Cookie.CookieError
	}
	val, ok := store.Values["key"].(string)
	if !ok {
		return "", public.Session.NoSession
	}
	return val, nil
}
