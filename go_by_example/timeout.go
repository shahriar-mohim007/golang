package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    // Simulate a long-running operation
    go func() {
        time.Sleep(3 * time.Second)
        ch <- "operation result"
    }()

    select {
    case res := <-ch:
        fmt.Println("Received:", res)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout! Operation took too long.")
    }
}
