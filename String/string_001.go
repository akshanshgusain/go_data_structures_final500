package String

func ReverseString(ip string) string {
	i := []rune(ip)
	li := 0
	ri := len(i) - 1
	for li <= ri {
		i[li], i[ri] = i[ri], i[li]
		li++
		ri--
	}

	return string(i)
}
