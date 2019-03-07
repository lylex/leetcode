package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	mapping := map[int]bool{}
	for _, v := range nums1 {
		mapping[v] = true
	}

	var ret []int
	for _, v := range nums2 {
		if mapping[v] {
			ret = append(ret, v)
			mapping[v] = false
		}
	}
	return ret
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
