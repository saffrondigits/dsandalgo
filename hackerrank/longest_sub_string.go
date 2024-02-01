// 1. Longest Increasing Subsequence

// A subsequence is created by deleting zero or more elements from a sequence while maintaining order.
// A sequence is strictly increasing if every element in the sequence is greater than the previous element. For example, [1, 2,
// 3] is strictly increasing while [1, 2, 2] is not.

// Given an array of integers, determine the length of the longest strictly increasing subsequence.

// For example, s = [1, 2, 2, 3, 1, 6]. Two strictly increasing subsequences are (1,2), (1, 2, 3).
// The longest increasing subsequence has a length of 4: LIS - [1,2,3,6].

// Simple input 0
// STDIN
// _______
// 3
// 1
// 4
// 3

// Sample output 0
// 2
// Explanation: The longest increasing subsequence [1,4]

// Simple input 1
// STDIN
// _______
// 5
// 1
// 4
// 5
// 2
// 6

// Sample output 1
// 4
// Explanation: The longest increasing subsequence [1,4,5,6]

// Simple input 1
// STDIN
// _______
// 4
// 2
// 3
// 3
// 5

// Sample output 1
// 3

// Explanation: The longest increasing subsequence [2,3,5]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func longestIncreasingSubsequence(arr []int) int {
	n := len(arr)
	lis := make([]int, n)

	for i := 0; i < n; i++ {
		lis[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}

	max := 0
	for _, v := range lis {
		if max < v {
			max = v
		}
	}

	return max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(longestIncreasingSubsequence(arr))
}
