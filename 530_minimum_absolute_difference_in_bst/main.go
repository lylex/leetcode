package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	midOrder(root.Right, arr)
}

func getMinimumDifference(root *TreeNode) int {
	var arr []int
	midOrder(root, &arr)

	fmt.Println(arr)

	var res int = math.MaxInt64
	for i := 1; i < len(arr); i++ {
		if diff := arr[i] - arr[i-1]; res > diff {
			res = diff
		}
	}
	return res
}

func main() {
	fmt.Println("hello world!")

	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	fmt.Println(getMinimumDifference(root))
}
