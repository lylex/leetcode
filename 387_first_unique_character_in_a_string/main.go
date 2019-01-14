package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	mapping := [26]int{}
	for _, v := range s {
		mapping[v-'a']++
	}

	for k, v := range s {
		if mapping[v-'a'] == 1 {
			return k
		}
	}

	return -1
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(firstUniqChar("loveleetcode"))
}
