package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {

	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {

	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for num, count := range counts {
		heap.Push(pq, &Item{value: num, priority: count})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(pq).(*Item).value
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k)) // Output: [1, 2]

	nums = []int{1}
	k = 1
	fmt.Println(topKFrequent(nums, k)) // Output: [1]
}
