package main

import (
	"fmt"
)

// 这里是一个非常简单粗暴的解法
func rotate0(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}

	k %= length
	for i := 0; i < k; i++ {
		temp := nums[length-1]
		for j := length - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = temp
	}
}

// 这里涉及到一个切片作为参数，如何修改元素的问题，直接修改了便是了
// 可以将切片理解为一个mallock出来的一个堆内存，即便复制了，这个堆内存还是共享的
func rotate(nums []int, k int) {
        reverse := func(a []int, start, end int) {
                for i, j := start, end; i < j; i, j = i+1, j-1 {
                        a[i], a[j] = a[j], a[i]
                }
        }

        length := len(nums)
        if length == 0 {
                return
        }
        k = k % length
        reverse(nums, 0, length-k-1)
        reverse(nums, length-k, length-1)
        reverse(nums, 0, length-1)

}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(input, 3)
	fmt.Println(input)
}
