// the file is the entry point for the API Gateway service.

package main

import (
	"log"
	"net/http"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081") //  we are reading the HTTP address from environment variables, if it does not exist, we default to ":8081"
)

// starting the server
func main() {
	log.Println("Starting API Gateway")

	/*
		http.HandleFunc --> It connects a URL path to a function that should run when someone visits that path. Think of it like telling your server: “Whenever someone visits THIS URL, run THAT function.”
		func(w http.ResponseWriter, r *http.Request) { ... } --> This is an anonymous function that takes two parameters: w (which is used to send responses back to the client) and r (which contains information about the incoming request).
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from API Gateway"))
	})

	http.ListenAndServe(httpAddr, nil) // this starts HTTP server on the specified address and "nil" means we are using the default ServeMux
}
