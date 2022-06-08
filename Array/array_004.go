package Array

// Array_004 Method 1: Simple counting : time : O(n), Space : O(3)
//func Array_004(ip []int) {
//	zc := 0
//	oc := 0
//	tc := 0
//
//	for _, e := range ip {
//		if e == 0 {
//			zc++
//		} else if e == 1 {
//			oc++
//		} else {
//			tc++
//		}
//	}
//
//	for i := 0; i < len(ip); i++ {
//		if i < zc {
//			ip[i] = 0
//		} else if i < oc+zc {
//			ip[i] = 1
//		} else {
//			ip[i] = 2
//		}
//	}
//}

// Array004 Array_004 Dutch National flag Approach
func Array004(ip []int) {
	low := 0
	mid := 0
	hi := len(ip) - 1

	for mid <= hi {
		switch ip[mid] {
		case 0:
			ip[low], ip[mid] = ip[mid], ip[low]
			low++
			mid++
			break
		case 1:
			mid++
			break
		case 2:
			ip[mid], ip[hi] = ip[hi], ip[mid]
			hi--
			break
		}
	}
}
