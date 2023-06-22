package main

import (
	"fmt"
	"suanfa.tree/tree"
)


func main() {
	var t tree.Tree
	t = tree.Tree{}
	t.Root = &tree.Node{Value: 2}
	t.Root.Left = &tree.Node{Value: 3}
	t.Root.Right = &tree.Node{Value: 4}
	t.Root.Left.Left = &tree.Node{Value: 5}
	t.Root.Left.Right = &tree.Node{Value: 6}
	res := inorderTraversal(t.Root)
	fmt.Println(res)

	// 层次遍历
	fmt.Println("层次遍历")
	t.SelectTree()
	fmt.Printf("\n")
	// 先序遍历
	fmt.Println("先序遍历")
	t.SelectFront(t.Root)
	fmt.Printf("\n")
	// 中序遍历
	fmt.Println("中序遍历")
	t.SelectMid(t.Root)
	fmt.Printf("\n")
	// 后序遍历
	fmt.Println("后序遍历")
	t.SelectBack(t.Root)
	fmt.Printf("\n")
}


func inorderTraversal(root *tree.Node) (res []int) {

	return traversal(root, &res)
}

func traversal(node *tree.Node, res *[]int) []int {
	if node == nil {
		return *res
	}
	traversal(node.Left, res)
	*res = append(*res, node.Value)
	traversal(node.Right, res)
	return *res
}