package main

import (
	"fmt"
)

func minMoves1(nums []int) int {
	length := len(nums)
	var round int
	for {
		max := nums[0]
		maxIndex := 0
		for i := 1; i < length; i++ {
			if max < nums[i] {
				max = nums[i]
				maxIndex = i
			}
		}
		if max == nums[0] {
			return round
		}

		round++
		for i := 0; i < length; i++ {
			if i == maxIndex {
				continue
			}
			nums[i]++
		}
	}
	return round
}

func minMoves(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	var sum int
	for _, v := range nums {
		sum += v - min
	}

	return sum
}

func main() {
	fmt.Println("hello world!")
	a := []int{1, 2, 3}
	fmt.Println(minMoves(a))
}
