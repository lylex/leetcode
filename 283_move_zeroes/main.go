package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			i++
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
}

func main() {
	fmt.Println("hello world!")
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
