package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}

func mergeList(l1, l2 *ListNode)  *ListNode{
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
}
