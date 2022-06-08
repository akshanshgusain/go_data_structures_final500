package Array

import "fmt"

// For Sorted Arrays
// 1. union, time: O(m+n)
func sortedUnion(arr1 []int, arr2 []int) {
	n := len(arr1)
	m := len(arr2)
	i := 0
	j := 0

	for i < n && j < m {
		if arr1[i] < arr2[j] {
			fmt.Printf(" %d << ", arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			fmt.Printf(" %d << ", arr2[i])
			j++
		} else {
			if n > m {
				fmt.Printf(" %d << ", arr1[i])
			} else {
				fmt.Printf(" %d << ", arr2[i])
			}

			i++
			j++
		}
	}
}

// 2. Intersection
func sortedIntersection(a1 []int, a2 []int) {
	n := len(a1)
	m := len(a2)
	i := 0
	j := 0
	fmt.Println()
	for i < n && j < m {
		if a1[i] < a2[j] {
			i++
		} else if a1[i] > a2[j] {
			j++
		} else {
			if n > m {
				fmt.Printf(" %d ", a1[i])
			} else {
				fmt.Printf(" %d ", a2[i])
			}

			i++
			j++
		}
	}
}

//For unSorted Arrays
// 1. Union
//Method 1: Use STL Set, time :O(mlog(m) + nlog(n))
//Method 2: Use STL Map, time :O(m+n)
// Method 3: Sort the arrays and apply union for sorted array,
// time:O(mlog(m) + nlog(n)).

func unSortedUnion(a1 []int, a2 []int) {
	fmt.Println()
	// Map as Set
	hm := make(map[int]bool)
	for _, e := range a1 {
		hm[e] = true
	}
	for _, e := range a2 {
		hm[e] = true
	}

	for k, v := range hm {
		if v {
			fmt.Printf(" %d ", k)
		}
	}
}

//2. Intersection
// Method 1: User STL Map.(If the elements in each array are unique)
// Method 2: User HashSet, time: O(m+n)
func unSortedIntersection(a1 []int, a2 []int) {

	fmt.Println()
	// Map as Set
	hm := make(map[int]bool)
	if len(a1) > len(a2) {
		a1, a2 = a2, a1 // better to iterate over a shorter set
	}

	for _, e := range a1 {
		hm[e] = true
	}
	for _, e := range a2 {
		if hm[e] {
			fmt.Printf(" %d ", e)
		}
	}

}

func Array006() {
	arr1 := []int{1, 2, 4, 5, 6, 7}
	arr2 := []int{2, 3, 5, 7}

	sortedUnion(arr1, arr2)
	sortedIntersection(arr1, arr2)

	arr3 := []int{3, 5, 1, 7, 9, 8}
	arr4 := []int{1, 51, 17, 95, 8}
	unSortedUnion(arr3, arr4)
	unSortedIntersection(arr3, arr4)
}
