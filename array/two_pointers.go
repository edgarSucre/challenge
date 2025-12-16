package array

import (
	"strings"
	"unicode"
)

func IsPalindromeS(s string) bool {
	s = strings.ToUpper(s)

	runes := make([]rune, 0, len(s))

	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			runes = append(runes, v)
		}
	}

	if len(runes) == 0 {
		return true
	}

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}
