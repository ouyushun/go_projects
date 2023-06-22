package main

import (
	"fmt"
	_ "fmt"
	_ "suanfa.chain/LinkList"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func main() {
	str := "一调度"
	for i, s := range str {
		fmt.Printf("%c， %v", s, i)
	}


	l1 := ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{3, nil}

	l2 := ListNode{4, nil}
	l2.Next = &ListNode{5, nil}
	l2.Next.Next = &ListNode{6, nil}

	res := addTwoNumbers(&l1, &l2)


	for {
		fmt.Println(res)
		res = res.Next
		if (res.Next == nil) {
			fmt.Println(res)
			break
		}
	}

}

/**
 * Definition for singly-linked list.

 */


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//虚拟head
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

