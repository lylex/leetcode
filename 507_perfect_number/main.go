package main

import (
	"fmt"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for i := 2; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum == num
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(checkPerfectNumber(28))
}
