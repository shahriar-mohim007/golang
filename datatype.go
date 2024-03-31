// Go provides both signed and unsigned integer arithmetic. There are four distinct sizes of
// signed integers—8, 16, 32, and 64 bits—represented by the types int8, int16, int32, and
// int64, and corresponding unsigned versions uint8, uint16, uint32, and uint64.

// Signed numbers are represented in 2’s-complement form, in which the high-order bit is
// reserved for the sign of the number and the range of values of an n-bit number is from −2n−1
// to 2n−1−1. Unsigned integers use the full range of bits for non-negative values and thus have
// the range 0 to 2n−1. For instance, the range of int8 is −128 to 127, whereas the range of
//uint8 is 0 to 255.
package main

import "fmt"

func main() {
    var u uint8 = 255
    fmt.Println(u, u+1, u*u)

    var i int8 = 127
    fmt.Println(i, i+1, i*i)
	var apples int32 = 1
    var oranges int16 = 2
	var compote = int(apples) + int(oranges)
	fmt.Println(compote)
	f := 3.141 // a float64
    integer := int(f) // Corrected variable name to avoid redeclaration
    fmt.Println(f, integer)
	// // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"
	s := "hello, world"
    fmt.Println(len(s))
    // "12"
    fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	var x float32 = math.Pi
    var y float64 = math.Pi
    var z complex128 = math.Pi
}