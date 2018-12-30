package main

import (
	"fmt"
)

// TreeNode is a definition.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSumWithCurrentSum(root *TreeNode, currentSum, sum int) int {
	if root == nil {
		return 0
	}
	currentSum += root.Val
	temp := 0
	if currentSum == sum {
		temp++
	}
	return temp + pathSumWithCurrentSum(root.Left, currentSum, sum) + pathSumWithCurrentSum(root.Right, currentSum, sum)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return pathSumWithCurrentSum(root, 0, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func main() {
	tree := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: -2,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}
	fmt.Println(pathSum(&tree, 8))

	fmt.Println("hello world!")
	tree2 := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
	fmt.Println(pathSum(&tree2, 3))
}
