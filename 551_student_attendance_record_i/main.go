package main

import (
	"fmt"
)

func checkRecord(s string) bool {
	var a, continousLCount int
	for _, v := range s {
		switch v {
		case 'A':
			a++
			continousLCount = 0
			if a > 1 {
				return false
			}
		case 'L':
			continousLCount++
			if continousLCount > 2 {
				return false
			}
		case 'P':
			continousLCount = 0
		}
	}
	return true
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
	fmt.Println(checkRecord("PPALLL"))
	fmt.Println(checkRecord("LLPPPLL"))
}
