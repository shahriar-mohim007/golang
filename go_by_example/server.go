package main

import (
    "fmt"
    "net/http"
)

// Handler function for the root path
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    // Register the handler function for the root path
    http.HandleFunc("/", helloHandler)

    // Start the server on port 8080
    fmt.Println("Server starting on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
