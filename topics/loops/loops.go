//the basic loop in Go is written in standard C-like syntax:

for INITIAL; CONDITION; AFTER{
  // do something
}

// INITIAL is run once at the beginning of the loop and can create
// variables within the scope of the loop.

// CONDITION is checked before each iteration. If the condition doesn't pass
// then the loop breaks.

// AFTER is run after each iteration.

// For example:

for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9
package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1 + (float64(i) * 0.01)
	}
	return totalCost
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}