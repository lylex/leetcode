package main

import (
	"fmt"
)

func isPowerOfFour1(num int) bool {
	if num < 1 {
		return false
	}

	for num != 1 {
		if num%4 != 0 {
			return false
		}
		num /= 4
	}
	return true
}

func isPowerOfFour2(num int) bool {
	if num < 1 {
		return false
	}

	var zeros int
	for num != 1 {
		if num&1 != 0 {
			return false
		}
		zeros++
		num >>= 1
	}
	return zeros%2 == 0
}

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}

	for num != 1 {
		if num&3 != 0 {
			return false
		}
		num >>= 2
	}

	return true
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(5))
}
