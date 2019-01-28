package main

import (
	"fmt"
)

func convertToBase7(num int) string {
	var ret []byte
	isPositive := true
	if num < 0 {
		isPositive = false
		num = -num
	}
	for num != 0 {
		ret = append(ret, byte('0'+num%7))
		num /= 7
	}
	if !isPositive {
		ret = append(ret, '-')
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	if len(ret) == 0 {
		ret = []byte{'0'}
	}
	return string(ret)
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
}
