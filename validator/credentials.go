package validator

import (
	"mailbase/shared/public"
	"mailbase/shared/templates"
	"net/http"
	"unicode"
)

func Password(w http.ResponseWriter, r *http.Request, Password string) {
	r.Close = true

	// Password check Why not regexp? Because re2 does not support lookaheads '?= '
	count, low, up, num := 0, false, false, false
	for _, s := range Password {
		if !unicode.IsLetter(s) && !unicode.IsNumber(s) {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, "Invalid password format: Only numbers/letters are allowed")
			return
		}
		switch {
		case unicode.IsLower(s):
			low = true
		case unicode.IsUpper(s):
			up = true
		case unicode.IsNumber(s):
			num = true
		}
		count++
	}

	if public.Register.PasswordMin > count || count > public.Register.PasswordMax {
		_ = templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Register.InvalidLength)
		return
	}
	if !(low && up && num) {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Register.InvalidFormat)
		return
	}

	r.Close = false
}
