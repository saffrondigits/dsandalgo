package main

import "fmt"

// Reverse a number 123 -> 321
// Reverse a number -456 -> -654

func main() {
	number := 42

	reverse := 0

	for number != 0 {
		rem := number % 10
		reverse = reverse*10 + rem
		number = number / 10
	}

	fmt.Printf("reverse: %v\n", reverse)
}
