package main

import (
	"fmt"
)

func compress(chars []byte) int {
	var current byte
	var count int
	var index int // 指向chars中需要修改的位置
	for k, v := range chars {
		count++
		if v != current || k == len(chars)-1 {
			if current != 0 {
				chars[index] = current
				index++
				if count > 1 {
					if k == len(chars)-1 {
						count++
					}
					chars[index] = byte(count - 1 + 48)
					index++
				}
				current = v
				count = 1
				continue
			}
			current = v
		}
	}

	//	chars = chars[:index] // not work
	//	fmt.Println(string(chars))
	// copy(chars, chars[:index]) // not work
	//	chars = append([]byte{}, chars[:index]...)

	temp := make([]byte, 0)
	copy(temp, chars[:index])
	chars = make([]byte, 0, 0)
	chars = append(chars, temp...)

	return index
}

func main() {
	fmt.Println("hello world!")
	a := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(a))
	fmt.Println(string(a))
}
