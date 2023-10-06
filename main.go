package main

import (
	"Final500Go/Heap"
	"container/heap"
	"fmt"
)
 
func main() {
	maxHeap := &Heap.MaxHeap{10, 5, 15, 8, 20}
	heap.Init(maxHeap)

	// Push elements into the max heap
	heap.Push(maxHeap, 25)
	heap.Push(maxHeap, 12)

	// Pop elements from the max heap (they will be in descending order)
	for maxHeap.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(maxHeap))
	}
}
