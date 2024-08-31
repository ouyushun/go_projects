package main

import "github.com/kr/pretty"

type Teacher struct {
	Name string
	Age int
}

func (t Teacher) ChangeName() {
	t.Name = "newName"
}
type Student struct {
	Name string
	Age int
}
func (s *Student) ChangeName() {
	s.Name = "newName"
}

func main() {
	t := Teacher{Name: "teacher1"}
	t.ChangeName()
	pretty.Println(t)

	s := Student{Name: "student1"}
	s.ChangeName()
	pretty.Println(s)
}
