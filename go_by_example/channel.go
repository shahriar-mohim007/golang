package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 42 // Send value to channel
    }()

    value := <-ch // Receive value from channel
    fmt.Println(value)
}
