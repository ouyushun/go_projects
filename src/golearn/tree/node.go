package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func main() {
	var root = treeNode{3, nil, nil}
	root.left = &treeNode{4, nil, nil}
	root.right = &treeNode{5, nil, nil}
	fmt.Println(root)

	nodes := []treeNode{
		{3, &treeNode{6, nil, nil}, nil},
	}
	fmt.Println(nodes)


}
