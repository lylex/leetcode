package main

import (
	"fmt"
)

func findPairs1(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	mapping := make(map[int]int)
	for _, v := range nums {
		mapping[v]++
	}

	var ret int
	for value, count := range mapping {
		if k == 0 {
			if count > 1 {
				ret += 2
			}
			continue
		}

		t1 := value - k
		t2 := value + k
		if mapping[t1] != 0 {
			ret++
		}
		if mapping[t2] != 0 {
			ret++
		}
	}
	return ret / 2
}

func findPairs(nums []int, k int) int {
	ans, m := 0, make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for n, v := range m {
		if k < 0 {
			return ans
		} else if k == 0 && v >= 2 {
			ans++
		} else if k > 0 && m[n+k] > 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
}
