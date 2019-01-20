package main

import (
	"fmt"
)

func constructRectangle(area int) []int {
	temp := (area + 1) / 2
	for temp*temp > area {
		temp /= 2
	}

	for temp*temp < area {
		temp++
	}

	for area%temp != 0 {
		temp--
	}

	max := temp
	if area/temp > max {
		max = area / temp
	}
	return []int{max, area / max}
}

func main() {
	fmt.Println("hello world!")
	fmt.Println(constructRectangle(5))
}
