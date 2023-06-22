package main

import "fmt"


func main() {
	l1 := ListNode{1, nil}
	l2 := ListNode{2, nil}

	l1.Next = &l2
	fmt.Println(*l1.Next)
}

func middleNode(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	slow, fast := newHead, newHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil {
		return slow
	}
	return slow.Next
}
