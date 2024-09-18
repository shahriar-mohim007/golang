package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
)

// Secret key for signing JWTs
var jwtKey = []byte("your_secret_key")

// GenerateJWT generates a new JWT token for the given username
func GenerateJWT(username string) (string, error) {
    // Create a new token object
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires in 1 hour

    // Sign the token with our secret key
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenStr string) (jwt.Claims, error) {
    // Parse the token
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        // Verify the token's signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, fmt.Errorf("invalid token")
    }
}
