package main

import "fmt"

// Find the most consecutive occurrences of a char in a string
// "aaabbccccccdeeeee"
// "cccccc"
func main() {
	r, c := mostConsecutiveOccurrence("aaabbccccccdeeeeeee")
	fmt.Printf("r: %v\n", r)
	fmt.Printf("c: %v\n", c)
}

func mostConsecutiveOccurrence(s string) (string, int) {
	var maxRune rune
	var maxCount int
	var lastRune rune
	var lastCount int

	for _, r := range s {
		if r == lastRune {
			lastCount++
		} else {
			lastRune = r
			lastCount = 1
		}

		if lastCount > maxCount {
			maxRune = lastRune
			maxCount = lastCount
		}
	}

	return string(maxRune), maxCount
}
