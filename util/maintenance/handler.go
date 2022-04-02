package maintenance

import (
	"github.com/illiafox/mailbase/shared/templates"
	"net/http"
)

const TimeDefault = "Unknown"

var (
	down  bool
	until = TimeDefault
	works = templates.CompileTemplate("../shared/templates/site/works/index.html")
)

func Off(time string) {
	down = true
	until = time
}

func Works() bool {
	return !down
}

func On() {
	down = false
	until = TimeDefault
}

type Handler struct {
	Root http.Handler
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if down {
		works.WriteAnyCode(w, http.StatusInternalServerError, until)
	} else {
		h.Root.ServeHTTP(w, r)
	}
}
