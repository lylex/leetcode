package main

import "fmt"

func isPalindrome(s string) bool {
	ret := true
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			ret = false
			break
		}
		i++
		j--
	}
	return ret
}

func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}

	var maxLen, left, right int
	for i := 0; i < len(s); i++ {
		for j := i + maxLen; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				maxLen = j - i + 1
				left = i
				right = j
			}
		}
	}
	return s[left : right+1]
}

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	var maxLen, left, right int
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				dp[j][i] = true
			}
			if i == j+1 {
				dp[j][i] = s[i] == s[j]
			}
			if i > j+1 {
				dp[j][i] = s[i] == s[j] && dp[j+1][i-1]
			}

			if dp[j][i] && i-j+1 > maxLen {
				maxLen = i - j + 1
				left = j
				right = i
			}
		}
	}
	return s[left : right+1]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
