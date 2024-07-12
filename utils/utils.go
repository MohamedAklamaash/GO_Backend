package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		log.Fatal("body cannot be empty")
	}
	return json.NewDecoder(r.Body).Decode(&payload)
}

func WriteJSON(w http.ResponseWriter, statuscode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, statuscode int, err error) {
	WriteJSON(w, statuscode, map[string]string{"error": err.Error()})
}
