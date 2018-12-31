package main

import (
	"fmt"
)

func numToString(n int) string {
	if n < 0 || n > 15 {
		return ""
	}

	if n < 10 {
		return string(byte('0' + n))
	}

	return string(byte(n - 10 + 'a'))
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	if num < 0 {
		num += -1&0xffffffff + 1
	}

	ret := ""
	for num != 0 {
		ret = numToString(num%16) + ret
		num >>= 4 // num /= 16 位移更有效率
	}
	return ret
}

// 解法二
func toHex2(num int) (answer string) {
	m := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = (1 << 32) + num
	}
	for num > 0 {
		answer = string(m[num%16]) + answer
		num >>= 4
	}
	return answer
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(toHex(26))
}
