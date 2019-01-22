package main

import (
	"fmt"
)

func findWords(words []string) []string {
	mapping := map[rune]int{
		'Q': 1,
		'W': 1,
		'E': 1,
		'R': 1,
		'T': 1,
		'Y': 1,
		'U': 1,
		'I': 1,
		'O': 1,
		'P': 1,

		'A': 2,
		'S': 2,
		'D': 2,
		'F': 2,
		'G': 2,
		'H': 2,
		'J': 2,
		'K': 2,
		'L': 2,

		'Z': 3,
		'X': 3,
		'C': 3,
		'V': 3,
		'N': 3,
		'B': 3,
		'M': 3,
	}

	ret := make([]string, 0)
	for _, w := range words {
		var row int
		match := true
		for _, v := range w {
			if v <= 'z' && v >= 'a' {
				v -= 32
			}
			if row == 0 {
				row = mapping[v]
				continue
			}
			if mapping[v] != row {
				match = false
			}
		}

		if match {
			ret = append(ret, w)
		}
	}
	return ret
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
