package String

import (
	"fmt"
	"unicode"
)

func isPalindromeHelper(ip string) bool {
	i := []rune(ip)
	li := 0
	ri := len(ip) - 1

	//unicode.ToLower(i[li])

	for li <= ri {
		if unicode.ToLower(i[li]) != unicode.ToLower(i[ri]) {
			return false
		}
		li++
		ri--
	}

	return true
}

func IsPalindrome(ip string) {
	if isPalindromeHelper(ip) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
