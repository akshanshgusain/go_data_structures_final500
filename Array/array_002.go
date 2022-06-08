package Array

func Array002(ip []int) (min, max int) {
	min = ip[0]
	max = ip[0]

	for _, v := range ip {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}
