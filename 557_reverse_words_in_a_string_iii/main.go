package main

import (
	"fmt"
)

func reverse(s []byte) {
	if len(s) == 0 {
		return
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return
}

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}

	ret := []byte(s)
	head := 0
	tail := 1
	for tail <= len(s) {
		if tail == len(s) {
			reverse(ret[head:])
			break
		}

		if ret[tail] != ' ' {
			tail++
			continue
		}

		reverse(ret[head:tail])
		head = tail + 1
		tail = head + 1
	}
	return string(ret)
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
