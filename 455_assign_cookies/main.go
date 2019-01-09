package main

import (
	"fmt"
)

func sort(g []int) {
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {
			if g[i] > g[j] {
				g[i], g[j] = g[j], g[i]
			}
		}
	}
}

func findContentChildren(g []int, s []int) int {
	sort(g)
	sort(s)

	var ret int
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			ret++
			i++
			j++
			continue
		}
		j++
	}

	return ret
}

func main() {
	fmt.Println("hello world!")
	g1 := []int{1, 2, 3}
	s1 := []int{1, 1}
	fmt.Println(findContentChildren(g1, s1))

	g2 := []int{1, 2}
	s2 := []int{1, 2, 3}
	fmt.Println(findContentChildren(g2, s2))
}
