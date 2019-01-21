package main

import (
	"fmt"
)

func nextGreaterElement(findNums []int, nums []int) []int {
	mapping := make(map[int]int)
	for k, v := range nums {
		mapping[v] = k
	}

	ret := make([]int, len(findNums))
	for k, v := range findNums {
		ret[k] = -1
		for i := mapping[v] + 1; i < len(nums); i++ {
			if nums[i] > v {
				ret[k] = nums[i]
				break
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
