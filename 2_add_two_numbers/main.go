package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	promote := 0
	var result, p *ListNode
	for l1 != nil || l2 != nil || promote != 0 {
		if result == nil {
			result = &ListNode{}
			p = result
		} else {
			p.Next = &ListNode{}
			p = p.Next
		}

		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Val = (sum + promote) % 10
		promote = (sum + promote) / 10
	}
	return result
}

func main() {
	fmt.Println("hello world!")

	left := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	right := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	result := addTwoNumbers(left, right)
	fmt.Println(result)
	fmt.Println(result.Next)
	fmt.Println(result.Next.Next)
}
