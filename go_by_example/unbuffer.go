package main

import (
    "fmt"
)

func main() {
    ch := make(chan int,2)

    go func() {
        ch <- 42 
		
		// This send operation will block until main() receives the value
    }()

    val := <-ch 
	val1 := <-ch// Receive the value sent by the goroutine
    fmt.Println(val)
	fmt.Println(val1)
}
