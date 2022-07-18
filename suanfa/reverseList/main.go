package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	a := []int{1,2,3,4}
	a = append(a, 5)





	head := ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	res := reverseList(&head)
	fmt.Println(res)
	fmt.Println(res.Next)
	fmt.Println(res.Next.Next)
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	pre := &ListNode{}
	for curr != nil{
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}