package main

import (
        "fmt"
)

func arrayPairSum(nums []int) int {
        for i:=0; i<len(nums); i++ {
                for j := i+1; j < len(nums); j++ {
                        if nums[i] > nums[j] {
                                nums[i], nums[j] = nums[j], nums[i]
                        }
                }
        }

        var sum int
        for i:=0; i< len(nums)/2 ; i++ {
                sum+=nums[2*i]
        }
        return sum
}

func main() {
        fmt.Println("hello world!")
        fmt.Println(arrayPairSum([]int{1,4,3,2}))
}
