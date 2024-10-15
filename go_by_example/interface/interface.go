package main

import (
	"fmt"
	"net/http"
)

// Define a function type that matches the Handler signature
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Implement the ServeHTTP method for the HandlerFunc type
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r) // Call the actual function
}

// MyHandler is a simple HTTP handler function
func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Use MyHandler as an HTTP handler
	http.Handle("/", HandlerFunc(MyHandler))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
