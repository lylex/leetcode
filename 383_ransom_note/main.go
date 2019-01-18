package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	mapping := map[rune]int{}
	for _, v := range ransomNote {
		mapping[v]++
	}

	total := len(ransomNote)
	for _, v := range magazine {
		if _, ok := mapping[v]; ok && mapping[v] != 0 {
			mapping[v]--
			total--
		}

		if total == 0 {
			break
		}
	}

	return total == 0
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
