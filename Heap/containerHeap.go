package Heap

import (
	"container/heap"
	"fmt"
)

// MaxHeap Helper function to convert a slice into a max-heap.
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // '>' for max-heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k <= 0 || k > n {
		return nil
	}

	// Initialize the max-heap.
	h := &MaxHeap{}
	heap.Init(h)

	var result []int

	// Build the initial window's heap.
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	// Process the rest of the elements.
	for i := k; i < n; i++ {
		result = append(result, (*h)[0])

		// Remove the leftmost element of the window from the heap.
		heap.Remove(h, 0)

		// Add the next element to the window heap.
		heap.Push(h, nums[i])
	}

	// Add the maximum value of the last window.
	result = append(result, (*h)[0])

	return result
}

func runner() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	result := maxSlidingWindow(nums, k)
	fmt.Println(result) // Output: [3 3 5 5 6 7]
}

func main() {
	maxHeap := &MaxHeap{10, 5, 15, 8, 20}
	heap.Init(maxHeap)

	// Push elements into the max heap
	heap.Push(maxHeap, 25)
	heap.Push(maxHeap, 12)

	// Pop elements from the max heap (they will be in descending order)
	for maxHeap.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(maxHeap))
	}
}
