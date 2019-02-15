package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func diameterOfBinaryTreeWithMax(root *TreeNode, max *int) {
	if root == nil {
		return
	}
	current := maxDepth(root.Left) + maxDepth(root.Right)
	if current > *max {
		*max = current
	}
	diameterOfBinaryTreeWithMax(root.Left, max)
	diameterOfBinaryTreeWithMax(root.Right, max)
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	diameterOfBinaryTreeWithMax(root, &max)
	return max
}

func main() {
	fmt.Println("hello world!")
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}
