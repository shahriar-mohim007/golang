package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Minute)
	fmt.Println("done")
}

//In your Go program, the time.Sleep(time.Minute)
//function is crucial for observing the output of the goroutines. Here's why:
//
//Goroutines and Concurrency: Goroutines are concurrent
//threads of execution. When you call go f("goroutine")
//and go func(msg string) { fmt.Println(msg) }("going"),
//these functions run asynchronously, meaning the
// main function doesn't wait for them to complete.
//
// Main Function Exit: If the main function
// completes before the goroutines,
// the program will exit, and the goroutines
// won't have a chance to run.
//The time.Sleep(time.Minute) call ensures that the main
//function doesn't exit immediately, giving the goroutines time to execute.
