// Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:

// // func make([]T, len, cap) []T
// mySlice := make([]int, 5, 10)

// // the capacity argument is usually omitted and defaults to the length
// mySlice := make([]int, 5)

// Slices created with make will be filled with the zero value of the type.

// If we want to create a slice with a specific set of values, we can use a slice literal:

// mySlice := []string{"I", "love", "go"}

// Note that the array brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.
// Length

// The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:

// mySlice := []string{"I", "love", "go"}
// fmt.Println(len(mySlice)) // 3

// Capacity

// The capacity of a slice is the number of elements in the underlying array,
//counting from the first element in the slice. It is accessed using the built-in cap() function:

// mySlice := []string{"I", "love", "go"}
// fmt.Println(cap(mySlice)) // 3

// Generally speaking, unless you're hyper-optimizing the memory usage of your program,
// you don't need to worry about the capacity of a slice because it will automatically grow as needed.

package main

func getMessageCosts(messages []string) []float64 {
	// Preallocate a slice for the costs with the same length as the messages slice
	costs := make([]float64, len(messages))

	// Calculate the cost for each message and store it in the corresponding index of the costs slice
	for i, message := range messages {
		costs[i] = float64(len(message)) * 0.01
	}

	return costs
}
