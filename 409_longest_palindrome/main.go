package main

import "fmt"

func longestPalindrome(s string) int {
	var mapping [255]int
	for _, v := range s {
		mapping[byte(v)]++
	}

	var hasSingle bool
	var pairs int
	for _, v := range mapping {
		pairs += v & 0xfffe
		if v%2 != 0 {
			hasSingle = true
		}
	}

	if hasSingle {
		pairs++
	}

	return pairs
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
