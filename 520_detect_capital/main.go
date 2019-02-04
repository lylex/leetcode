package main

import (
	"fmt"
)

func isCapital(r rune) bool {
	return r <= 'Z' && r >= 'A'
}

func detectCapitalUse1(word string) bool {
	if len(word) == 1 {
		return true
	}

	var leftAllCapital bool
	if isCapital(rune(word[0])) {
		// 首字母大写，则全大写或全小写
		leftAllCapital = isCapital(rune(word[1]))
	} else {
		// 首字母小写，则全小写
		leftAllCapital = false
	}

	for i := 1; i < len(word); i++ {
		if isCapital(rune(word[i])) != leftAllCapital {
			return false
		}
	}

	return true
}

func detectCapitalUse(word string) bool {
	var count int
	for _, v := range word {
		if isCapital(v) {
			count++
		}
	}
	return count == 0 || count == len(word) || count == 1 && isCapital(rune(word[0]))
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("Google"))
	fmt.Println(detectCapitalUse("FlaG"))
}
