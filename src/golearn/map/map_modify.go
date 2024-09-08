package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Printf("map的类型%T\n", m)

	map_modify()
}
func map_modify() {
	users := make(map[int]string)
	users[1] = "user1"

	fmt.Printf("before modify: userdd:%v\n", users[1])  // before modify: userdd:user1
	modify(users)
	fmt.Printf("after modify: userdd:%v\n", users[1])  // after modify: userdd:user2
}

func modify(u map[int]string) {
	u[1] = "user2"
}

