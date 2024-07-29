package main

import (
	"fmt"
	"net/http"
)

// Define a function type
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Implement the http.Handler interface
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

// Define a simple HTTP handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Use the function as an HTTP handler
	http.Handle("/hello", HandlerFunc(helloHandler))

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
