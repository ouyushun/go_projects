package LinkList
//结点定义如下。
type Node struct {
	Data int
	Next *Node
}

//链表定义如下。其中仅包含一个头结点，表明这是一个单链表。
type List struct {
	Length int
	Head   *Node
}

type LinkList struct {
	val int
	Next   *LinkList
}


//链表初始化函数
//链表初始化即创建一个头结点，并返回指向链表对象的指针。
func InitList() *List {
	L := new(List)
	L.Head = new(Node)
	return L
}

func (list *List) IsNull() bool {
	if list.Length > 0 {
		return false
	}
	return true
}
//链表的插入函数-头插
//在链表头插入一个结点，同时为了保证代码的鲁棒性，考虑用户不调用初始化函数就直接调用头插的情况。
func (list *List) PushFront(v int) {
	node := &Node{Data: v}
	if list.IsNull() {
		list.Head = node
		list.Length++
		return
	}
	node.Next = list.Head
	list.Head = node
	list.Length++
	return
}

//链表的插入函数-尾插
//在链表尾部插入一个函数，每次需要从头结点开始遍历，直到找到链表尾，再在链表尾部插入新结点。可以考虑存储一个链表尾结点，来提高尾插的效率。
func (list *List) PushBack(v int) {
	node := &Node{Data: v}
	if list.IsNull() {
		list.Head.Next = node
	} else {
		cur := list.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
	list.Length++
	return
}

