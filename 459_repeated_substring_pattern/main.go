package main

import (
	"fmt"
)

func isContains(a []int, n int) bool {
	for _, v := range a {
		if n == v {
			return true
		}
	}
	return false
}

func repeatedSubstringPattern1(s string) bool {
	mapping := map[rune][]int{}
	for k, v := range s {
		if !isContains(mapping[v], k) {
			mapping[v] = append(mapping[v], k)
		}
	}

	minLen := 0 // 找到一个出现次数最少的字母，以及它出现的次数
	var minKey rune
	for k, v := range mapping {
		if minLen == 0 || minLen > len(v) {
			minLen = len(v)
			minKey = k
		}
	}

	// 如果被选出的字母，只出现了一次的话，肯定不是了
	if minLen < 2 {
		return false
	}

	// 把这个选出的最小出现次数的字母作为模板，看看它是否符合等差数列，如果不符合，这个字符串肯定不符合
	list := mapping[minKey]
	diff := (list[minLen-1] - list[0]) / (minLen - 1)
	for i := 1; i < minLen; i++ {
		if list[i] != list[i-1]+diff {
			return false
		}
	}

	for key, candidates := range mapping {
		if minKey == key {
			continue
		}

		if len(candidates)%minLen != 0 {
			return false
		}

		for i := 0; i < len(candidates)/minLen; i++ {
			for j := i + diff; j < len(candidates); j += len(candidates) / minLen {
				if candidates[j] != candidates[j-len(candidates)/minLen]+diff {
					return false
				}
			}
		}
	}
	return true
}

func repeatedSubstringPattern(s string) bool {
	for i := len(s) / 2; i >= 1; i-- {
		if len(s)%i != 0 {
			continue
		}
		temp := ""
		for j := 0; j < len(s)/i; j++ {
			temp += s[:i]
		}
		if temp == s {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
