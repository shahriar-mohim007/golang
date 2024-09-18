package main

import (
    "fmt"
    "net/http"
)
func loginHandler(w http.ResponseWriter, r *http.Request) {
    var user map[string]string
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil || user["username"] == "" || user["password"] == "" {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // In a real application, you should validate the username and password from a database
    if user["username"] != "testuser" || user["password"] != "password" {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Generate JWT
    token, err := GenerateJWT(user["username"])
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Return the token
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the token from the Authorization header
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        // Validate the token
        _, err := ValidateJWT(tokenStr)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Token is valid, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is a protected endpoint")
}

func main() {
    http.HandleFunc("/login", loginHandler)

    // Protect the /protected route with the authentication middleware
    http.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
