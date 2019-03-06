package main

import (
	"fmt"
)

const HightestPowerOfThree = 1162261467

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	return HightestPowerOfThree%n == 0
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(isPowerOfThree(9))

}
