package main

import "fmt"

// Remove duplicates from an array
func main() {
	names := []string{
		"Dave",
		"Kalidasu",
		"Rakshit",
		"Bob",
		"Tom",
		"Dave",
	}

	fmt.Printf("removeDuplicate: %v\n", removeDuplicate(names))
}

func removeDuplicate(str []string) []string {
	seen := make(map[string]bool)

	result := []string{}

	for _, s := range str {
		_, ok := seen[s]
		if !ok {
			seen[s] = true
			result = append(result, s)
		}
	}

	return result
}
