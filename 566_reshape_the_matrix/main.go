package main

import (
        "fmt"
)

func matrixReshape1(nums [][]int, r int, c int) [][]int {
	if nums == nil || len(nums) == r || len(nums)*len(nums[0]) != r*c {
		return nums
	}

	arr := make([]int, r*c, r *c)

	row := len(nums)
	col := len(nums[0])
	for i:= 0; i < row; i++ {
		for j := 0; j < col; j++ {
			arr[i*col + j] = nums[i][j]
		}
	}

	ret := make([][]int, r, r)
	for i := 0; i < r; i++ {
		temp := make([]int, c, c)
		for j := 0; j < c; j++ {
			temp[j] = arr[i*c + j]
		}
		ret[i] = temp
	}
	return ret
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if nums == nil || len(nums) == r || len(nums)*len(nums[0]) != r*c {
		return nums
	}

	ret := make([][]int, 0)
	for i:= 0; i < r; i++ {
		ret = append(ret, make([]int, c, c))
	}

	col := len(nums[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			total := i*c+j
			ret[i][j] = nums[total/col][total%col]
		}
	}
		
	return ret
}

func main() {
	nums := [][]int{
		[]int{1,2},
		[]int{3,4},
	}
        fmt.Println(matrixReshape(nums, 1, 4))
        fmt.Println(matrixReshape(nums, 2, 3))
        fmt.Println("hello world!")
}
