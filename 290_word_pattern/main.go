package main

import (
	"fmt"
	"strings"
)

func wordPattern1(pattern string, str string) bool {
	contain := strings.Split(str, " ")
	length := len(pattern)
	if length != len(contain) {
		return false
	}

	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if pattern[i] == pattern[j] {
				if strings.Compare(contain[i], contain[j]) != 0 {
					return false
				}
			} else {
				if strings.Compare(contain[i], contain[j]) == 0 {
					return false
				}
			}
		}
	}

	return true
}

func wordPattern(pattern string, str string) bool {
	mapping := map[rune][]int{}
	for k, v := range pattern {
		if mapping[v] == nil {
			mapping[v] = []int{}
		}
		mapping[v] = append(mapping[v], k)
	}

	candidates := strings.Split(str, " ")
	diffs := []string{}

	// 组内相等
	for _, list := range mapping {
		fmt.Println(diffs)
		diffs = append(diffs, candidates[list[0]])
		for i := 1; i < len(list); i++ {
			if strings.Compare(candidates[list[i]], candidates[list[0]]) != 0 {
				return false
			}
		}
	}

	// 组外不等
	for i := 0; i < len(diffs); i++ {
		for j := i + 1; j < len(diffs); j++ {
			if strings.Compare(diffs[i], diffs[j]) == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
