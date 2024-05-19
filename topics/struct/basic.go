// We use structs in Go to represent structured data. It's often convenient to group different types
// of variables together. For example, if we want to represent a car we could do the following:

// type car struct {
//   maker string
//   model string
//   doors int
//   mileage int
// }

// This creates a new struct type called car. All cars have a maker, model, doors and mileage.

// Structs in Go are often used to represent data that you might use a dictionary or object for in other languages.
package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

// don't edit below this line

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
}
