package main

import (
	"fmt"
)

func licenseKeyFormatting(S string, K int) string {
	var ret []byte
	var count int
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}

		if count == K {
			count = 0
			ret = append(ret, byte('-'))
		}

		count++
		c := S[i]
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		ret = append(ret, byte(c))

	}

	if len(ret) > 1 {
		for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
			ret[i], ret[j] = ret[j], ret[i]
		}
	}
	return string(ret)

}

func main() {
	fmt.Println("hello world!")
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
}
