package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// 以t为根的树的所有的节点和
func sum(t *TreeNode) int {
	if t == nil {
		return 0
	}
	return t.Val + sum(t.Left) + sum(t.Right)
}

// t节点的坡度
func tilt(t *TreeNode) int {
	if t == nil {
		return 0
	}

	return absDiff(sum(t.Left), sum(t.Right))
}

func findTilt1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return tilt(root) + findTilt(root.Left) + findTilt(root.Right)
}

func postOrder(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	leftSum := postOrder(root.Left, res)
	rightSum := postOrder(root.Right, res)

	*res += absDiff(leftSum, rightSum)
	return leftSum + rightSum + root.Val
}

func findTilt(root *TreeNode) int {
	var res int
	postOrder(root, &res)
	return res
}

func main() {
	fmt.Println("hello world!")
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(findTilt(tree))
}
