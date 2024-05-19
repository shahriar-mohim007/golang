// Functions in Go can take zero or more arguments.

// To make code easier to read, the variable type comes after the variable name.

// For example, the following function:

// func sub(x int, y int) int {
//   return x-y
// }

// Accepts two integer parameters and returns another integer.

// Here, func sub(x int, y int) int is known as the "function signature".

package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// don't touch below this line

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
