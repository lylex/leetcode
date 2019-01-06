package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for k, v := range nums {
		mapping[v] = k
	}

	for k, v := range nums {
		if x, ok := mapping[target-v]; ok {
			return []int{k, x}
		}
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {
	list := map[int]int{}

	if len(nums) < 2 {
		return nil
	}

	for k, v := range nums {
		left := target - v
		if vv, ok := list[left]; ok {
			return []int{vv, k}
		}

		list[v] = k
	}

	return nil
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
