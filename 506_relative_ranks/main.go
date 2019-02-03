package main

import (
	"fmt"
)

func findRelativeRanks(nums []int) []string {
	ret := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		greaterCount := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] > nums[i] {
				greaterCount++
			}
		}

		var out string
		switch greaterCount {
		case 0:
			out = "Gold Medal"
		case 1:
			out = "Silver Medal"
		case 2:
			out = "Bronze Medal"
		default:
			out = fmt.Sprintf("%d", greaterCount+1)
		}

		ret[i] = out
	}

	return ret
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
}
