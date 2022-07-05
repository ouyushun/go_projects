package tree


type Node2 struct {
	Value interface{}
	Left Node
	Right Node
}

type tree2 struct {
	root *Node2
}