package main

import (
	"fmt"
)

func findComplement(num int) int {
	temp := num
	var last int
	for temp != 0 {
		last = temp
		temp &= temp - 1
	}
	return last<<1 - 1 - num
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
}
