package main

import (
	"fmt"
)

func isPalindrome(n int) bool {
	var str []int
	for n != 0 {
		str = append(str, n%10)
		n /= 10
	}
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func biggestN(n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret *= 10
	}
	return ret - 1
}

func smallestN(n int) int {
	if n <= 1 {
		return 0
	}

	ret := 1
	for i := 0; i < n-1; i++ {
		ret *= 10
	}
	return ret
}

func largestPalindrome(n int) int {
	for i := biggestN(n); i >= smallestN(n); i-- {
		for j := i; j >= smallestN(n); j-- {
			product := i * j
			if isPalindrome(product) {
				return product % 1337
			}
		}
	}
	return 9
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(largestPalindrome(3))
	fmt.Println(isPalindrome(131))
}
