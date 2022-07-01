package main

import (
	"fmt"
	""
)
func main() {
	list := LinkList.InitList()
	// 添加链表结点
	fmt.Println("添加结点")
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.Insert(3, 4)
	list.ShowList()
	// 按索引删除链表结点
	fmt.Println("删除索引为1的结点")
	list.Delete(1)
	list.ShowList()
	// 按值删除链表结点
	fmt.Println("删除值为1的结点")
	list.Remove(1)


}
