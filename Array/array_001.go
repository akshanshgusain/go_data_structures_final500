package Array

import (
	"fmt"
)

func Array001(ip []int) {
	n := len(ip)

	for i := 0; i < n/2; i++ {
		ip[i], ip[n-i-1] = ip[n-i-1], ip[i]
	}

	for _, e := range ip {
		fmt.Printf("%d\t", e)
	}
}
