// MinhindenSuketcing istrart) take the array
// of strings stored in stales, which will contain only two strings, the first parameter being the string N and the second parameter being a string K of some characters, and your goal is to determine the smallest substring of N that contains all the characters in K. For example: if strAns is ['aaabaaddae", "aed' then the smallest substring of N that contains the characters a, e, and d is "dae" located at the end of the string. So for this example your program should return the string dae.
// Another example: if strar is l'aabdecdbeaed",
// "and' then the smallest substring of N that contains all of the characters in K is "aabd" which is located at the beginning of the string. Both parameters will be strings ranging in length from 1 to 50 characters and all of Ks characters will exist somewhere in the string N. Both strings will only contains lowercase alphabetic characters.

package main

import (
	"fmt"
	"math"
)

// MinWindowSubstring returns the smallest substring of N that contains all the characters in K
func MinWindowSubstring(strArr []string) string {
	// get the input strings
	N := strArr[0]
	K := strArr[1]

	// initialize the hash maps for K and the current window
	needToFind := make(map[byte]int)
	hasFound := make(map[byte]int)
	for i := 0; i < len(K); i++ {
		needToFind[K[i]]++
		hasFound[K[i]] = 0
	}

	// initialize the pointers and the count of matched characters
	start := 0
	end := 0
	count := 0

	// initialize the result variables
	minLength := math.MaxInt32
	minStart := 0
	minEnd := 0

	// loop through the string N
	for end < len(N) {
		// get the current character
		ch := N[end]

		// if the character is not in K, skip it
		if _, ok := needToFind[ch]; !ok {
			end++
			continue
		}

		// update the hash map and the count for the current character
		hasFound[ch]++
		if hasFound[ch] <= needToFind[ch] {
			count++
		}

		// if all the characters in K are matched, try to shrink the window
		if count == len(K) {
			// advance the start pointer as long as the window is still valid
			for start <= end {
				// get the character at the start pointer
				ch := N[start]

				// if the character is not in K, skip it
				if _, ok := needToFind[ch]; !ok {
					start++
					continue
				}

				// update the result variables if the current window is smaller
				if end-start+1 < minLength {
					minLength = end - start + 1
					minStart = start
					minEnd = end
				}

				// update the hash map and the count for the character at the start pointer
				hasFound[ch]--
				if hasFound[ch] < needToFind[ch] {
					count--
					start++
					break
				}

				// advance the start pointer
				start++
			}
		}

		// advance the end pointer
		end++
	}

	// return the smallest substring or an empty string if none exists
	if minLength == math.MaxInt32 {
		return ""
	} else {
		return N[minStart : minEnd+1]
	}
}

func main() {
	// test the function
	fmt.Println(MinWindowSubstring([]string{"aaabaaddae", "aed"}))          // dae
	fmt.Println(MinWindowSubstring([]string{"aabdccdbcacd", "aad"}))        // aabd
	fmt.Println(MinWindowSubstring([]string{"ahffaksfajeeubsne", "jefaa"})) // aksfaje
}
