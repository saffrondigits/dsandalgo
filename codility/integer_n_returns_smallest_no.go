package solution

import (
	"math"
	"strconv"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	if N == 1 {
		return 0
	}
	numDigits := len(strconv.Itoa(N))
	return int(math.Pow10(numDigits - 1))
}
