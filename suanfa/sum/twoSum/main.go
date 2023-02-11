package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = &ListNode{}

}


func TwoSum(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := new(ListNode)
	curr := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1  = l1.Next
		}
		if l2 != nil {
			carry +=  l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry = carry / 10
		curr.Next = new(ListNode)
		curr = curr.Next
	}

	return head
}