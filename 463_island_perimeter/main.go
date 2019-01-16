package main

import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
	ilength := len(grid)
	jlength := len(grid[0])
	sum := 0
	for i := 0; i < ilength; i++ {
		for j := 0; j < jlength; j++ {
			if grid[i][j] == 0 {
				continue
			}

			// 边际
			if i == 0 {
				sum++
			}
			if i == ilength-1 {
				sum++
			}
			if j == 0 {
				sum++
			}
			if j == jlength-1 {
				sum++
			}

			if j-1 >= 0 {
				if grid[i][j-1] == 0 { // left
					sum++
				}
			}
			if j+1 < jlength {
				if grid[i][j+1] == 0 { // right
					sum++
				}
			}
			if i-1 >= 0 {
				if grid[i-1][j] == 0 { // up
					sum++
				}
			}
			if i+1 < ilength {
				if grid[i+1][j] == 0 { // down
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	input := [][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
	}

	fmt.Println("hello world!")
	fmt.Println(islandPerimeter(input))
}
