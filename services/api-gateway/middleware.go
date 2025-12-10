package main

import "net/http"

// basically this middleware enables CORS for the given HTTP handler function, this middleware will be called before the actual handler function is executed
func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// if the request method is OPTIONS, we just return with 200 OK status, i.e. allow pre-flight request from broswer API.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler(w, r) // calling the actual handler function
	}
}
