package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // Create a channel to receive signals
    sigs := make(chan os.Signal, 1)

    // Notify the channel of SIGINT and SIGTERM signals
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // Create a channel to signal when the program should exit
    done := make(chan bool, 1)

    // Start a goroutine to handle signals
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println("Received signal:", sig)

        // Perform cleanup here if needed

        done <- true
    }()

    // Simulate work by sleeping for 10 seconds
    fmt.Println("Press Ctrl+C to exit")
    time.Sleep(10 * time.Second)

    // Wait for the signal handler to finish
    <-done
    fmt.Println("Exiting gracefully")
}
