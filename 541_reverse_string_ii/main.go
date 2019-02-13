package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	arr := []byte(s)
	var round int
	for ; round < len(arr)/k; round++ {
		if round%2 != 0 {
			continue
		}

		for i, j := round*k, (round+1)*k-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// handle overflow uncompleted round
	if len(arr)%k != 0 && len(arr)/k%2 == 0 {
		for i, j := len(arr)/k*k, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	return string(arr)
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcdefg", 8))
}
