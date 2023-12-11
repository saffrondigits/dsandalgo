package main

import (
	"fmt"
	"strings"
)

// Check for palindrome
// Radar
// A man, a plan, a canal, Panama
func main() {
	isPal := isPalindrome("Radar")
	fmt.Printf("is palindrome: %v\n", isPal)
}

func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	r := reverseStr([]byte(s))

	if s == r {
		return true
	}
	return false
}

func reverseStr(s []byte) string {
	j := len(s) - 1

	for i := 0; i < len(s); i++ {
		if i < j {
			temp := s[i]
			s[i] = s[j]
			s[j] = temp
		}
		j--
	}

	return string(s)
}

// eat
// tea
// ate

// aet
// aet
// aet
