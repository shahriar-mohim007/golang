// Arrays are fixed-size groups of variables of the same type.

// The type [n]T is an array of n values of type T

// To declare an array of 10 integers:

// var myInts [10]int

// or to declare an initialized literal:

// primes := [6]int{2, 3, 5, 7, 11, 13}
package main

import "fmt"

const (
	retry1 = "click here to sign up"
	retry2 = "pretty please click here"
	retry3 = "we beg you to sign up"
)

func getMessageWithRetries() [3]string {
	return [3]string{
		retry1,
		retry2,
		retry3,
	}
}

// don't touch below this line

func send(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("complete failure")
		}
	}
}

func main() {
	send("Bob", 0)
	send("Alice", 1)
	send("Mangalam", 2)
	send("Ozgur", 3)
}
