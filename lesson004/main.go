package main

import "fmt"

// Factorial of a number

// Fact 5 = 5 x 4 x 3 x 2 x 1
// Fact 3 = 3 x 2 x 1

func main() {
	result := factorial(5)
	fmt.Printf("result: %v\n", result)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	fact := 1

	for i := 1; i <= n; i++ {
		fact = fact * i
	}

	return fact
}
