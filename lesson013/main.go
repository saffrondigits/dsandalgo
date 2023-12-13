package main

import (
	"fmt"
	"math"
)

// Finding maximum and minimum in an array
func main() {
	arr := []int{4, 1, 2, 3, 4, 5, 6}

	min, max := findMinMax(arr)
	fmt.Printf("min: %v\n", min)
	fmt.Printf("max: %v\n", max)

}

func findMinMax(arr []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt

	for _, num := range arr {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	return min, max
}
