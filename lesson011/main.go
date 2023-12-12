package main

import "fmt"

func main() {
	fmt.Println(isAnagram("listen", "silent"))
}

// listen, silent
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := map[rune]int{}

	for _, r := range s1 {
		count[r]++
	}

	fmt.Printf("count 1: %v\n", count)
	for _, r := range s2 {
		count[r]--
		if count[r] < 0 {
			return false
		}
	}

	return true
}
