package main

import (
	"fmt"
)

// 解法一
//func str2int(str string) int {
//	var ret int
//	for _, n := range str {
//		ret = ret*10 + int(byte(n)-48)
//	}
//	return ret
//}
//
//func int2str(n int) string {
//	ret := make([]byte, 0)
//	for n != 0 {
//		x := n % 10
//		n /= 10
//		ret = append([]byte{byte(x) + 48}, ret...)
//	}
//	if len(ret) == 0 {
//		return "0"
//	}
//	return string(ret)
//}
//
//func addStrings(num1 string, num2 string) string {
//	return int2str(str2int(num1) + str2int(num2))
//}

func reverseStr(str string) string {
	s := []byte{}
	for _, n := range str {
		s = append([]byte{byte(n)}, s...)
	}
	return string(s)
}

func addStrings(num1 string, num2 string) string {
	num1, num2 = reverseStr(num1), reverseStr(num2)
	maxLen := len(num1)
	if maxLen < len(num2) {
		maxLen = len(num2)
	}

	isPromote := false
	var ret []byte
	for i := 0; i < maxLen; i++ {
		var left, right byte
		if i < len(num1) {
			left = byte(num1[i]) - 48
		}
		if i < len(num2) {
			right = byte(num2[i]) - 48
		}
		sum := left + right
		if isPromote {
			sum++
		}
		ret = append(ret, sum%10+48)
		isPromote = sum/10 == 1
	}
	// handle the last promote
	if isPromote {
		ret = append(ret, 49)
	}

	return reverseStr(string(ret))
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(addStrings("12", "29"))
}
