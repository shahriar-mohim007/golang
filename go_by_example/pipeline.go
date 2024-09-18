package main

import (
	"fmt"
)

func main() {
	// Create channels to pass data between goroutines
	naturals := make(chan int)
	squares := make(chan int)

	// Counter Goroutine
	// This goroutine sends an increasing sequence of integers to the `naturals` channel.
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer Goroutine
	// This goroutine receives integers from the `naturals` channel, squares them, and sends the results to the `squares` channel.
	go func() {
		for {
			x := <-naturals // Receive a value from the `naturals` channel
			squares <- x * x // Square the value and send it to the `squares` channel
		}
	}()

	// Printer (in main goroutine)
	// This loop in the main goroutine continuously receives and prints squared values from the `squares` channel.
	for {
		fmt.Println(<-squares) // Print the squared value received from the `squares` channel
	}
}
