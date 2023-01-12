package core

import (
	"encoding/json"
	"net/http"
)

func EncodeJsonResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}
