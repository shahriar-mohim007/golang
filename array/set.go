package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	// Create a set from nums
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	maxi := 0

	// Iterate through the set
	for num := range numSet {
		// Check if it's the start of a sequence
		if _, found := numSet[num-1]; !found {
			next := num + 1
			length := 1
			// Find the length of the sequence
			for {
				if _, found := numSet[next]; !found {
					break
				}
				next++
				length++
			}
			if length > maxi {
				maxi = length
			}
		}
	}

	return maxi
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // Output: 4
}
