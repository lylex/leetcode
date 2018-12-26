package main

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := 1; i <= n; i++ {
		var r string
		switch {
		case i%3 == 0 && i%5 == 0: // better to be i%15 == 0
			r = "FizzBuzz"
		case i%3 == 0:
			r = "Fizz"
		case i%5 == 0:
			r = "Buzz"
		default:
			r = fmt.Sprintf("%d", i)
		}
		ret[i-1] = r
	}
	return ret
}

func main() {
	fmt.Println(fizzBuzz(15))
}
