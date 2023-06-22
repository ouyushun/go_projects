package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	//res := reverseList(&head)
	res := doOp(&head)
	fmt.Println(res)
	fmt.Println(res.Next)
	fmt.Println(res.Next.Next)
}



func doOp(l *ListNode) *ListNode {
	var head *ListNode
	for l != nil {
		pre := head

		head = new(ListNode)
		head.Val = l.Val

		head.Next = pre



		l = l.Next
	}
	return head
}







func reverseList(head *ListNode) *ListNode {
	curr := head
	newLink := &ListNode{}
	for curr != nil{
		next := curr.Next
		curr.Next = newLink
		newLink = curr

		curr = next
	}
	return newLink
}

func Recerse2(head *ListNode) *ListNode {
	curr := head
	newHead := &ListNode{}

	for curr != nil {



		curr = curr.Next
	}
	return newHead
}
