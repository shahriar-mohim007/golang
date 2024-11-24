package main

import "fmt"

func main() {
	// Create the channel
	messages := make(chan string)

	// Send and receive message using the function
	go func() {
		// Send the message to the channel
		messages <- "ping"
		close(messages)
	}()

	// Receive the message and print it
	msg := <-messages
	fmt.Println(msg)
}
