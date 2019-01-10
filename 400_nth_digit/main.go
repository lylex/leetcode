package main

import (
	"fmt"
)

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	digits := 1 // 几位数，如1，10，100，1000
	start := 1  // 起始数，1, 10, 100, 1000
	width := 9  // 该位数有多少个数，9，90，900，9000
	sum := 0
	for {
		if n < sum+digits*width {
			targetNumber := (n-sum+1)/digits + start - 1 // 目标数
			targetIndex := (n-sum-1)%digits + 1          // 是目标数的左数第几位，1开始
			for i := 1; ; i++ {
				left := targetNumber % 10
				if i == digits-targetIndex+1 {
					return left
				}
				targetNumber /= 10
			}
		}
		sum += digits * width
		digits++
		start *= 10
		width *= 10
	}

	return -1
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(findNthDigit(11))
}
