package tree

import "fmt"

// 定义节点属性
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// 二叉树
type Tree struct {
	Root *Node
}

// 生成二叉树
func (tree *Tree) CreateTree(value int) {
	node := &Node{Value: value}
	fmt.Println("增加节点", value)
	if tree.Root == nil {
		tree.Root = node
		return
	}else {
		queue := []*Node{tree.Root}
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil {
				cur.Left = node
				return
			}else if cur.Right == nil {
				cur.Right = node
				return
			}else {
				queue = append(queue, cur.Left)
				queue = append(queue, cur.Right)
			}
		}
	}
}

// 层次遍历：即是广度优先方式搜索，使用队列方式实现
func (tree *Tree) SelectTree() {
	if tree.Root == nil {
		fmt.Println("not Root")
		return
	}
	var queue []*Node = []*Node{tree.Root}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Printf("%v ", cur.Value)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}

// 先序遍历：即是深度优先方式搜索，使用递归方式实现
// 根节点 --> 左节点 --> 右节点
func (tree *Tree) SelectFront(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.Value)
	tree.SelectFront(node.Left)
	tree.SelectFront(node.Right)
}

// 中序遍历：即是深度优先方式搜索，使用递归方式实现
// 左节点 --> 根节点 --> 右节点
func (tree *Tree) SelectMid(node *Node) {
	if node == nil {
		return
	}
	tree.SelectMid(node.Left)
	fmt.Printf("%v ", node.Value)
	tree.SelectMid(node.Right)
}

// 后序遍历：即是深度优先方式搜索，使用递归方式实现
// 左节点 --> 右节点 --> 根节点
func (tree *Tree) SelectBack(node *Node) {
	if node == nil {
		return
	}
	tree.SelectBack(node.Left)
	tree.SelectBack(node.Right)
	fmt.Printf("%v ", node.Value)
}