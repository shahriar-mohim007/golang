package main

import "fmt"

func main() {
	// declare here
	congrats := "happy birthday!"

	fmt.Println(congrats)
}

// Inside a function (like the main function) the := short assignment statement can be used in place of a var declaration.
// The := operator infers the type of the new variable based on the value. It's colloquially called the walrus operator because it looks like a walrus... sort of.

//Outside of a function (in the global/package scope), every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
