package main

import (
	"container/heap"
	"fmt"
)

type minIntHeap []int

func (h minIntHeap) Len() int           { return len(h) }
func (h minIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minIntHeap) Pop() interface{} {
	heapDerefrenced := *h

	size := len(heapDerefrenced)
	val := heapDerefrenced[size-1]
	*h = heapDerefrenced[:size-1]

	return val
}

func main() {
	h := &minIntHeap{}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 2)
	fmt.Printf("Min heap: %v\n", *h)
	fmt.Printf("Popped minimum: %v\n", heap.Pop(h))
}
