package main

import (
	"fmt"
)

func arrangeCoins(n int) int {
	if n < 2 {
		return n
	}

	f := func(k int) int {
		return (k+1)*k/2 - n
	}

	begin, end := 1, n
	for begin <= end {
		mid := begin + (end-begin)/2
		temp := f(mid)
		if temp == 0 {
			return mid
		} else if temp < 0 {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return begin - 1
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(arrangeCoins(8))
}
