package Array

// Array005 Method 1: Using Dutch National FLag
//func Array005(ip []int) {
//	lo := 0
//	hi := 0
//
//	for hi < len(ip) {
//		if ip[hi] >= 0 {
//			ip[lo], ip[hi] = ip[hi], ip[lo]
//			lo++
//			hi++
//		} else {
//			hi++
//		}
//	}
//}

// Array005 Method 2: Using 2 pointer Technique, O(n) time and O(1) Space
func Array005(ip []int) {
	i := 0
	j := len(ip) - 1

	for i <= j {
		if ip[i] < 0 && ip[j] < 0 {
			i++
		} else if ip[i] > 0 && ip[j] < 0 {
			ip[i], ip[j] = ip[j], ip[i]

		} else if ip[i] < 0 && ip[j] > 0 {
			j--
		} else if ip[i] > 0 && ip[j] > 0 {
			i++
		}
	}
}
