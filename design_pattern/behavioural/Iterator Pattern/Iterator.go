package main

import "fmt"

// Iterator interface defines methods to traverse a collection
type Iterator interface {
	HasNext() bool
	Next() int
}

// IntCollection represents a collection of integers
type IntCollection struct {
	numbers []int
}

// CreateIterator creates an iterator for the IntCollection
func (c *IntCollection) CreateIterator() Iterator {
	return &IntIterator{
		collection: c,
		index:      0,
	}
}

// IntIterator is the concrete iterator for the IntCollection
type IntIterator struct {
	collection *IntCollection
	index      int
}

// HasNext checks if there are more elements in the collection
func (it *IntIterator) HasNext() bool {
	return it.index < len(it.collection.numbers)
}

// Next returns the next element and advances the index
func (it *IntIterator) Next() int {
	if it.HasNext() {
		value := it.collection.numbers[it.index]
		it.index++
		return value
	}
	return 0 // or handle as an error
}

func main() {
	// Create an IntCollection
	collection := &IntCollection{
		numbers: []int{10, 20, 30, 40, 50},
	}

	// Create an iterator for the collection
	iterator := collection.CreateIterator()

	// Use the iterator to traverse the collection
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
