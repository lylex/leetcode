package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(n *TreeNode) bool {
	return n != nil &&
		n.Left == nil &&
		n.Right == nil
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if isLeaf(root.Left) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func main() {
	fmt.Println("hello world!")

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(sumOfLeftLeaves(root))
}
