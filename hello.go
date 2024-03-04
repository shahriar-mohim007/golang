package main

import "fmt"

func main() {
	/*
					go run hello.go
					The go run command does in fact compile your code into a binary.
					However, the binary is built in a temporary directory. The go run command
					builds the binary, executes the binary from that temporary directory, and
					then deletes the binary after your program finishes. This makes the go run
					command useful for testing out small programs or using Go like a scripting
					language.
					Go natively handles Unicode, so it can process text in all the world’s languages
					If the program is more than a one-shot experiment, it’s likely that you would want to compile
					it once and save the compiled result for later use. That is done with go build:
					This creates an executable binary file called helloworld that can be run any time without fur-
					ther processing:

					Let’s now talk about the program itself. Go code is organized into packages, which are similar
			to libraries or modules in other languages. A package consists of one or more .go source files
			in a single directory that define what the package does. Each source file begins with a package
			declaration, here package main, that states which package the file belongs to, followed by a list
			of other packages that it imports, and then the declarations of the program that are stored in
			that file.

		Package main is special. It defines a standalone executable program, not a library. Within
		package main the function main is also special—it’s where execution of the program begins.
		Whatever main does is what the program does. Of course, main will normally call upon func-
		tions in other packages to do much of the work, such as the function fmt.Println.
	*/
	fmt.Print("Hello World\n")
}
