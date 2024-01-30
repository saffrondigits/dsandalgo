package main

import "fmt"

func main() {
	end := 4

	for i := 1; i < end; i++ {
		fmt.Printf("%v", i)
		defer printNum(i)
	}
}

func printNum(n int) {
	fmt.Printf("%v", n)
}
