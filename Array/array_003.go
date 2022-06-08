package Array

import (
	"sort"
)

func Array003(ip []int, k int) (min, max int) {
	sort.SliceStable(ip, func(i, j int) bool {
		return ip[i] < ip[j]
	})

	return ip[k-1], ip[len(ip)-1-k]
}

// method 1: nLogN - heap/Merge sort and return k-1th element
// method 2: n + kLogN - create a min heap and do k pop operation
// method 3:  - Use an STL ordered Map

//method 4: Quick Select : Average case complexity: n, worst case : n2
// Quickselect is a selection algorithm to find the k-th smallest element in an unordered list.
