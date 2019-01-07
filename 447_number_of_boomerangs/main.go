package main

import (
	"fmt"
)

func numberOfBoomerangs(points [][]int) int {
	ret := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if j == i {
				continue
			}
			for k := 0; k < len(points); k++ {
				if k == i || k == j {
					continue
				}

				if (points[i][0]-points[j][0])*(points[i][0]-points[j][0])+(points[i][1]-points[j][1])*(points[i][1]-points[j][1]) ==
					(points[i][0]-points[k][0])*(points[i][0]-points[k][0])+(points[i][1]-points[k][1])*(points[i][1]-points[k][1]) {
					ret++
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("hello world!")
	points := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
	}
	fmt.Println(numberOfBoomerangs(points))
}
