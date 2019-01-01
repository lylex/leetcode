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

func findAnagrams(s string, p string) []int {
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

func main() {
	fmt.Println("hello world!")
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
