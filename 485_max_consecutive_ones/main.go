package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	var max, current int
	for _, v := range nums {
		if v == 1 {
			current++
		} else {
			current = 0
		}
		if current > max {
			max = current
		}
	}
	return max
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
