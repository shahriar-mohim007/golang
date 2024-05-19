// Named Return Values

// Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function.

// Named return values are best thought of as a way to document the purpose of the returned values.

// According to the tour of go:

//     A return statement without arguments returns the named return values.
// 	This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

// func getCoords() (x, y int){
//   // x and y are initialized with zero values

//   return // automatically returns x and y
// }

// Is the same as:

// func getCoords() (int, int){
//   var x int
//   var y int
//   return x, y
// }

// In the first example, x and y are the return values. At the end of the function,
// we could simply write return to return the values of those two variables, rather than writing return x,y.

package main

import (
	"fmt"
)

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

// don't edit below this line

func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====================================")
}

func main() {
	test(4)
	test(10)
	test(22)
	test(35)
}
