package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Read[T any](r *http.Request) (T, error) {
	var dst T
	err := json.NewDecoder(r.Body).Decode(&dst)
	return dst, err
}
