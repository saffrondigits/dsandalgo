package main

import "fmt"

func main() {

	fmt.Println(fib(10)) // prints 55
}

func fib(n int) int {
	// base case: n is 0 or 1, return n
	if n == 0 || n == 1 {
		return n
	}
	// recursive case: n is greater than 1, return fib(n-1) + fib(n-2)
	return fib(n-1) + fib(n-2)
}

// 1. Direct Recursion
// 2. Indirect Recursion
// 3. Tail Recursion
// 4. Linear Recursion
// 5. Tree Recursion
