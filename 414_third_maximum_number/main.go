package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	min -= 1

	first, second, third := min, min, min
	for _, n := range nums {
		switch {
		case n > first:
			third = second
			second = first
			first = n
		case n == first:
			// do nothing
		case n > second:
			third = second
			second = n
		case n == second:
			// do nothing
		case n > third:
			third = n
		}
	}

	if third == min {
		return first
	}
	return third
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(thirdMax([]int{1, 2, -2147483648}))
	fmt.Println(math.MinInt64)
}
