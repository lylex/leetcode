package main

import "fmt"

func lengthOfLongestSubstring1(s string) int {
	var max int
	for i := 0; i < len(s); i++ {
		mapping := make(map[byte]struct{}, len(s))
		for j := i; j < len(s); j++ {
			if _, ok := mapping[s[j]]; ok {
				break
			}
			tempMax := j - i + 1
			if tempMax > max {
				max = tempMax
			}
			mapping[s[j]] = struct{}{}
		}
	}
	return max
}

func lengthOfLongestSubstring(s string) int {
	mapping := make([]int, 256, 256)
	// set all to -1
	for i := 0; i < 255; i++ {
		mapping[i] = -1
	}

	var ret int
	left := -1
	for i := 0; i < len(s); i++ {
		// 如果当前值出现过，且在滑动窗口内
		if mapping[s[i]] != -1 && left < mapping[s[i]] {
			left = mapping[s[i]]
		}
		currentMax := i - left
		if currentMax > ret {
			ret = currentMax
		}
		mapping[s[i]] = i
	}
	return ret
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

}

