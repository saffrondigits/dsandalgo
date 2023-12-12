package main

import "fmt"

func main() {
	arr := []int{1, 3, 3, 3, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1}

	fmt.Println(mostFrequent(arr))
}

func mostFrequent(arr []int) int {
	counts := make(map[int]int)
	maxCount := 0

	maxValue := arr[0]

	for _, v := range arr {
		counts[v]++

		if counts[v] > maxCount {
			maxCount = counts[v]
			maxValue = v
		}
	}

	return maxValue
}

// 1, 3, 3, 3, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1

// counts[1:6, 3:3: 2:5]
