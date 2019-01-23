package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var mapping = make(map[int]int)
	findModeCore(root, mapping)

	var maxCount int
	for _, v := range mapping {
		if v > maxCount {
			maxCount = v
		}
	}

	var ret []int
	for k, v := range mapping {
		if v == maxCount {
			ret = append(ret, k)
		}
	}
	return ret
}

func findModeCore1(root *TreeNode, mapping map[int]int) {
	if root == nil {
		return
	}
	mapping[root.Val]++
	findModeCore(root.Left, mapping)
	findModeCore(root.Right, mapping)
	return
}

func main() {
	fmt.Println("hello world!")

	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	fmt.Println(findMode(tree))
}
