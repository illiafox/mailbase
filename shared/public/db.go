package public

import (
	"errors"
	"fmt"
	"net/http"
)

// InternalWithError is used to detect internal error instead of non-panic
type InternalWithError struct {
	InternalError error
}

func (int InternalWithError) Error() string {
	return int.InternalError.Error()
}

var ErrorInternal = errors.New("internal service error, try again later")

func NewInternalWithError(err error) InternalWithError {
	return InternalWithError{err}
}

func WriteWithCode(w http.ResponseWriter, statusCode int, elements ...any) (int, error) {
	w.WriteHeader(statusCode)
	return fmt.Fprintln(w, elements...)
}
