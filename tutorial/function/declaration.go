//Developers often wonder why the declaration syntax in Go is different from the tradition established in the C family of languages.
// C-Style syntax

// The C language describes types with an expression including the name to be declared, and states what type that expression will have.

// int y;

// The code above declares y as an int. In general, the type goes on the left and the expression on the right.

// Interestingly, the creators of the Go language agreed that the C-style of declaring types in signatures gets confusing really fast - take a look at this nightmare.

// int (*fp)(int (*ff)(int x, int y), int b)

// Go-style syntax

// Go's declarations are clear, you just read them left to right, just like you would in English.

// x int
// p *int
// a [3]int

// It's nice for more complex signatures, it makes them easier to read.

// f func(func(int,int) int, int) int
