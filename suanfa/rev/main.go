package main
import "fmt"

type Node struct {
	Val int
	Next *Node
}

func main() {
	l := new(Node)
	l = &Node{1, nil}
	l.Next = &Node{2, nil}
	l.Next.Next = &Node{3, nil}

	list := Reverse(l)

	fmt.Println(list.Val)
	fmt.Println(list.Next.Val)
	fmt.Println(list.Next.Next.Val)
}

func t (a *[]int) {

	fmt.Println(a)
}

func Reverse(l *Node) *Node {
	var head *Node
	for l != nil {
		pre := head
		//新节点
		head = new(Node)
		head.Val = l.Val
		//新节点的Next
		head.Next = pre
		l = l.Next
	}

	return head
}