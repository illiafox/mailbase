package public

import (
	"encoding/json"
	"net/http"
)

type errorJSON struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

var JSON = jsonTools{}

type jsonTools struct {
}

func (jsonTools) WriteErrorString(w http.ResponseWriter, err string, code ...int) error {
	for _, status := range code {
		w.WriteHeader(status)
	}
	return json.NewEncoder(w).Encode(errorJSON{Error: err, Ok: false})
}

func (jsonTools) WriteError(w http.ResponseWriter, err error, code ...int) error {
	for _, status := range code {
		w.WriteHeader(status)
	}
	return json.NewEncoder(w).Encode(errorJSON{Error: err.Error(), Ok: false})
}
