package main

import "fmt"

// Swap two variables using a temp variable
func main() {
	a := "Apple"
	b := "Orange"

	temp := a
	a = b
	b = temp

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}
