// A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: _

// For example:

// func getPoint() (x int, y int) {
//   return 3, 4
// }

// // ignore y value
// x, _ := getPoint()

// Even though getPoint() returns two values, we can capture the first one and ignore the second.

package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe"
}
