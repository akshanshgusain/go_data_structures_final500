package Array

import "fmt"

func rotateKTimes(ip []int, k int) {
	n := len(ip)
	k = k % n

	for i := 0; i < n; i++ {
		if i < k {
			fmt.Printf("%v ", ip[n-k+i])
		} else {
			fmt.Printf("%v ", ip[i-k])
		}
	}
}

func Array007(ip []int, k int) {
	rotateKTimes(ip, k)
}
