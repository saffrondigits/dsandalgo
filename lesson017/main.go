package main

import "fmt"

func main() {
	fmt.Printf("factorial: %v\n", factorial(5))
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
