package main

import "fmt"

// Bubble sort
func main() {
	arr := []int{4, 3, 6, 2, 7, 1, 5}
	result := sort(arr)
	fmt.Printf("result: %v\n", result)
}

func sort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {

			if arr[j] > arr[i] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}

	return arr
}
