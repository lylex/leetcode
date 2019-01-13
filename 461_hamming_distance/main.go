package main

import (
	"fmt"
)

func hammingDistance1(x int, y int) int {
	var ret int
	for x != 0 || y != 0 {
		if x&1 != y&1 {
			ret++
		}
		x >>= 1
		y >>= 1
	}
	return ret
}

func hammingDistance2(x int, y int) int {
	var ret int
	xor := x ^ y
	for xor != 0 {
		if xor&1 != 0 {
			ret++
		}
		xor >>= 1
	}
	return ret
}

func hammingDistance(x int, y int) int {
	var ret int
	xor := x ^ y
	for xor != 0 {
		ret++
		xor &= xor - 1
	}
	return ret
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(hammingDistance(1, 4))
}
