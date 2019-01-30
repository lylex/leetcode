package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	return getSum(a^b, (a&b)<<1)
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-2, 3))
}
