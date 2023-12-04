package main

import "fmt"

// Linear search
func main() {
	arr := []int{2, 4, 5, 6, 8, 20}

	search := 20

	for i := 0; i < len(arr); i++ {

		if search == arr[i] {
			fmt.Printf("searched: %v at index %v\n", search, i)
			return
		}
	}

	fmt.Printf("Oppss! cannot find the value\n")
}
