package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    // Create a new HTTP client
    client := &http.Client{}

    // Create a new GET request with custom headers
    req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
    if err != nil {
        log.Fatalf("Error creating request: %v", err)
    }

    req.Header.Set("User-Agent", "GoHttpClient/1.0")

    // Send the request
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error sending request: %v", err)
    }
    defer resp.Body.Close()

    // Read and print the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %v", err)
    }

    fmt.Println("Response Status:", resp.Status)
    fmt.Println("Response Body:", string(body))
}
