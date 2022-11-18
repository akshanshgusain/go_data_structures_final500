package Heap

import "fmt"

type Heap struct {
	array []int
}

func (h *Heap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Heapify method to rearrange elements in the array, Bottom to Top
func (h *Heap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Heapify Top to Bottom
func (h *Heap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//loop while the index has at least one child
	for l <= lastIndex {
		if l == lastIndex {
			// when left child is the only child
		} else if h.array[l] > h.array[r] {
			// when left child is greater than right child
			childToCompare = l
		} else {
			// when right child is larger
			childToCompare = r
		}
		// compare the array value of the current index to larger child and swap
		// is smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
		}
	}
}

// Extract and removes the largest key and removes it from the heap
func (h *Heap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("cannot extract from an empty array")
		return -999
	}
	if len(h.array) == 1 {
		return extracted
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyUp(0)
	return extracted
}

// get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get left child index
func left(i int) int {
	return 2*i + 1
}

// get right child index
func right(i int) int {
	return 2*i + 2
}

// swap two indices
func (h *Heap) swap(i int, i1 int) {
	h.array[i], h.array[i1] = h.array[i1], h.array[i]
}

func Execute() {
	m := &Heap{}
	buildHeap := []int{10, 20, 30}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	fmt.Println(m)
	fmt.Println("Running Extract()")
	m.Extract()
	fmt.Println(m)
	m.Extract()
	fmt.Println(m)
	m.Extract()
	fmt.Println(m)

}
