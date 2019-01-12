package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	mapping := [255]byte{}
	for _, v := range s {
		mapping[byte(v)]++
	}

	for _, v := range t {
		key := byte(v)
		if mapping[key] == 0 {
			return key
		}
		mapping[key]--
	}

	return 0
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(findTheDifference("abcd", "abcde"))
}
