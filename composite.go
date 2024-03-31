package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	// array of 3 integers
	fmt.Println(a[0])            // print the first element
	fmt.Println(a[len(a)-1])     // print the last element, a[2]

	Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2, 0} // Corrected array assignment to have 3 elements
	fmt.Println(r[2])

	q = [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

	symbols := [...]string{USD: "$", EUR: "€", GBP: "!", RMB: "¥"}

	// Example usage:
	fmt.Println("Symbol for USD:", symbols[USD])
	fmt.Println("Symbol for EUR:", symbols[EUR])
	fmt.Println("Symbol for GBP:", symbols[GBP])
	fmt.Println("Symbol for RMB:", symbols[RMB])

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

	months := [...]string{1: "January", /* ... */, 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	// ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]
	for _, s := range summer {
		for _, q := range Q2 {
		if s == q {
		fmt.Printf("%s appears in both\n", s)
		}
		}
		}
		// Using […] makes an array. Using [] makes a slice.
		
		var x []int
        x = append(x, 10)
		x = append(x, 5, 6, 7)

		y := make([]int, 5)
		y = append(x, 10)
}
