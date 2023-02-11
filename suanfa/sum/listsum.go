package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := ListNode{}
	l1.Val = 1
	l1.Next= &ListNode{Val: 2}
	fmt.Println(l1.Next.Val)

	addTwoNumbers(&l1, &l1)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head
	carry := 0

	for (l1 != nil || l2 != nil || carry > 0) {
		curr.Next = new(ListNode)
		curr = curr.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry /= 10
	}
	return head.Next
}