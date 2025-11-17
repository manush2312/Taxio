package main

// ! json.go prevents duplication and handles all JSON responses cleanly.

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data) // encoding the data into JSON and writing it to the response writer
}
