package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBSTWorker(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}

	convertBSTWorker(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	convertBSTWorker(root.Left, sum)

	return root
}

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	convertBSTWorker(root, &sum)
	return root
}

func main() {
	fmt.Println("hello world!")

	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 13,
		},
	}
	t := convertBST(root)

	fmt.Println(t.Val)
	fmt.Println(t.Left.Val)
	fmt.Println(t.Right.Val)
}
