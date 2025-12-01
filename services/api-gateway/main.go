// the file is the entry point for the API Gateway service.

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081") //  we are reading the HTTP address from environment variables, if it does not exist, we default to ":8081"
)

// starting the server
func main() {
	log.Println("Starting API Gateway")

	/*
		! Creating a custom HTTP multiplexer (router) instead of using the default one.
		! Reasons:
			! 1. Customization: Allows for more control over routing behavior.
			! 2. Middleware Integration: Easier to add middleware for logging, authentication, etc.
			! 3. Scalability: Better suited for complex applications with many routes.
	*/
	mux := http.NewServeMux() //  creating a new HTTP request multiplexer (router)

	/*
		! http.HandleFunc --> It connects a URL path to a function that should run when someone visits that path. Think of it like telling your server: “Whenever someone visits THIS URL, run THAT function.”
		! func(w http.ResponseWriter, r *http.Request) { ... } --> This is an anonymous function that takes two parameters: w (which is used to send responses back to the client) and r (which contains information about the incoming request).
	*/
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("Hello from API Gateway"))
	// })

	// http.ListenAndServe(httpAddr, nil) // ! this starts HTTP server on the specified address and "nil" means we are using the default ServeMux

	mux.HandleFunc("POST /trip/preview", handleTripPreview) // registering the trip preview handler with the custom multiplexer
	// ! web socket handlers
	mux.HandleFunc("/ws/drivers", handleDriversWebSocket)
	mux.HandleFunc("/ws/riders", handleRidersWebSocket)

	/*
		Why create an http.Server instead of http.ListenAndServe?

		Because with http.Server, you can later add:
		✔️ Timeouts
		✔️ TLS support
		✔️ Graceful shutdown
		✔️ Read/write timeouts
		✔️ Max header size
		✔️ Idle connection timeout
	*/
	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	// if err := server.ListenAndServe(); err != nil {
	// 	log.Printf("HTTP server error: %v", err)
	// }

	//! implementing graceful shutdown
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("server(HTTP) listening on %s", httpAddr)
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1) // Whenever the process receives SIGINT or SIGTERM, send that signal into the shutdown channel
	// ! os.Interrupt --> this signal is sent when the user types Ctrl+C
	// ! syscall.SIGTERM --> this signal is sent by Kubernetes when it wants to terminate the application
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM) // this is allows us to listen the signals that are coming outside of our application

	select {
	case err := <-serverErrors:
		log.Printf("Error starting the server: %v", err)

	case sig := <-shutdown:
		log.Printf("Server is shutting down due to signal: %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// server you have up to 10 seconds to finish any in-flight requests and close cleanly. If you don’t finish in that time, we’ll force-close.
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not gracefully shutdown the server: %v", err)
			server.Close()
		}

	}
}
