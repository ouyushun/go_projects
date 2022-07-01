package LinkList
//结点定义如下。
type Node struct {
	data int
	next *Node
}

//链表定义如下。其中仅包含一个头结点，表明这是一个单链表。
type List struct {
	length   int
	head *Node
}


//链表初始化函数
//链表初始化即创建一个头结点，并返回指向链表对象的指针。
func InitList() *List {
	L := new(List)
	L.head = new(Node)
	return L
}

func (list *List) IsNull() bool {
	if list.length > 0 {
		return true
	}
	return false
}
//链表的插入函数-头插
//在链表头插入一个结点，同时为了保证代码的鲁棒性，考虑用户不调用初始化函数就直接调用头插的情况。
func (list *List) PushFront(v int) {
	node := &Node{data: v}
	if list.IsNull() {
		list.head = node
		list.length++
		return
	}
	node.next = list.head
	list.head = node
	list.length++
	return
}

//链表的插入函数-尾插
//在链表尾部插入一个函数，每次需要从头结点开始遍历，直到找到链表尾，再在链表尾部插入新结点。可以考虑存储一个链表尾结点，来提高尾插的效率。
func (list *List) PushBack(v int) {
	node := &Node{data: v}
	if list.IsNull() {
		list.head.next = node
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
	list.length++
	return
}

