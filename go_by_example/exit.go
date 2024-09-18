package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("!")
	print("mohim")

	//os.Exit(0)

	os.Exit(3)
	//os.Exit(code) from the os package.
	//The code is an integer value where 0 typically indicates success, a
	//nd any non-zero value indicates an error or failure.

	//defers will not be run when using os.Exit,
	//so this fmt.Println will never be called.
}
