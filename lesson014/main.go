package main

import "fmt"

// Move all the 0's in an array at the end
// 1, 0, 2, 0, 3, 0, 4, 5
// 1, 2, 3, 4, 5, 0, 0, 0

func main() {
	arr := []int{1, 0, 2, 0, 3, 0, 4, 5, 0, 7}
	fmt.Printf("moveZerosToEnd: %v\n", moveZerosToEnd(arr))
}

func moveZerosToEnd(arr []int) []int {
	count := 0

	for i := 0; i < len(arr); i++ {

		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	for count < len(arr) {
		arr[count] = 0
		count++
	}

	return arr
}
