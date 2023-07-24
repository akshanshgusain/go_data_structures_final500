package String

import (
	"fmt"
	"unicode"
)

func CountDuplicates(ip string) {
	i := []rune(ip)
	m := make(map[rune]int)
	for _, ch := range i {
		m[unicode.ToLower(ch)]++
	}

	for k, v := range m {
		if v > 1 {
			fmt.Printf("%c : %v\n", k, v)
		}

	}
}
