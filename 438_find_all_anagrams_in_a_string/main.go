package main

import (
	"fmt"
)

func isMatch(a string, originMapping map[rune]int, count int) bool {
	if count == 0 {
		return false
	}

	if len(a) != count {
		return false
	}

	mapping := map[rune]int{}
	for k, v := range originMapping {
		mapping[k] = v
	}

	for _, n := range a {
		if mapping[n] == 0 {
			return false
		}
		mapping[n]--
	}

	return true
}

func findAnagrams2(s string, p string) []int {
	ret := []int{}
	if len(s) < len(p) {
		return ret
	}

	mapping := map[rune]int{}
	count := 0
	for _, n := range p {
		mapping[n]++
		count++
	}

	for i := 0; i <= len(s)-len(p); i++ {
		if isMatch(s[i:i+len(p)], mapping, count) {
			ret = append(ret, i)
		}
	}
	return ret
}

func findAnagrams(s string, p string) []int {
	ret := []int{}
	mapping := make([]int, 255)
	count := len(p)

	for _, n := range p {
		mapping[byte(n)]++
	}

	var left, right int
	for right < len(s) {
		if mapping[byte(s[right])] > 0 {
			count--
		}
		mapping[byte(s[right])]--

		if count == 0 {
			ret = append(ret, left)
		}

		// 区间长度等于p的长度的时候，扩大区间
		if right-left == len(p)-1 {
			// 如果左区边界属于p，则在其自增前，把count加一
			if mapping[byte(s[left])] >= 0 {
				count++
			}
			mapping[byte(s[left])]++
			left++
		}

		right++
	}

	return ret
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
