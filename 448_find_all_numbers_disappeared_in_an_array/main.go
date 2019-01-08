package main

import (
	"fmt"
)

func findDisappearedNumbers1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := nums[i]
		if index < 0 {
			index = -index
		}
		if nums[index-1] > 0 {
			nums[index-1] = -nums[index-1]
		}
	}

	var ret []int
	for k, v := range nums {
		if v > 0 {
			ret = append(ret, k+1)
		}
	}
	return ret
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	var ret []int
	for k, v := range nums {
		if v != k+1 {
			ret = append(ret, k+1)
		}
	}
	return ret
}

func main() {
	fmt.Println("hello world!")
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
