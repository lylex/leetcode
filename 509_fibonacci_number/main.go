package main

import (
	"fmt"
)

func fib1(N int) int {
	if N < 2 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

func fib(N int) int {
	if N < 2 {
		return N
	}

	first := 0
	second := 1
	for i := 2; i <= N; i++ {
		first, second = second, first+second
	}

	return second
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
}
