// Slices in Go

// 99 times out of 100 you will use a slice instead of an array when working with ordered lists.

// Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.

// A slice is a dynamically-sized, flexible view of the elements of an array.

// The zero value of slice is nil.

// Slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:

// primes := [6]int{2, 3, 5, 7, 11, 13}
// mySlice := primes[1:4]
// // mySlice = {3, 5, 7}

// The syntax is:

// arrayname[lowIndex:highIndex]
// arrayname[lowIndex:]
// arrayname[:highIndex]
// arrayname[:]

// Where lowIndex is inclusive and highIndex is exclusive.

// lowIndex, highIndex, or both can be omitted to use the entire array on that side of the colon.

package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}

// don't touch below this line

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func test(name string, doneAt int, plan string) {
	defer fmt.Println("=====================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("no response")
		}
	}
}

func main() {
	test("Ozgur", 3, planFree)
	test("Jeff", 3, planPro)
	test("Sally", 2, planPro)
	test("Sally", 3, "no plan")
}

// Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.
// Except for items with explicit dimensions such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

// Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.
// If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller,
// analogous to passing a pointer to the underlying array. A Read function can therefore accept a slice argument rather
// than a pointer and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the
// Read() method of the File type in package os:

// Referenced from Effective Go

// func (f *File) Read(buf []byte) (n int, err error)
