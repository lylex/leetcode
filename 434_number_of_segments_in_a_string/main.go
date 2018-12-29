package main

import (
	"fmt"
)

func countSegments(s string) int {
	ret := 0
	segmentBegin := false
	for _, n := range s {
		if segmentBegin {
			if n == ' ' {
				segmentBegin = false
			}
		} else {
			if n != ' ' {
				segmentBegin = true
				ret++
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(countSegments("Hello, my name is John"))
}
