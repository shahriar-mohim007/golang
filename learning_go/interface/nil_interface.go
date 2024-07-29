package main

import "fmt"

type Incrementer interface {
	Increment()
}

type Counter struct{}

func (c *Counter) Increment() {}

func main() {
	var pointerCounter *Counter
	fmt.Println(pointerCounter == nil) // prints true

	var incrementer Incrementer
	fmt.Println(incrementer == nil) // prints true

	incrementer = pointerCounter
	fmt.Println(incrementer == nil) // prints false
}
