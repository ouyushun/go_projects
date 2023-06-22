package main

import (
	"fmt"
	"sort"
)

type Int32List []int32

func (s Int32List) Len() int           { return len(s) }
func (s Int32List) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int32List) Less(i, j int) bool { return s[i] < s[j] }

type MySlice []int

func (x MySlice) Len() int           { return len(x) }
func (x MySlice) Less(i, j int) bool { return x[i] < x[j] }
func (x MySlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	s1 := []int {5, 3, 4, 8, 2, 9}

	mm := []int(s1)

	slice := MySlice(s1)
	fmt.Printf("%T", mm)
	sort.Sort(slice)



	s2 := sort.IntSlice(s1)
	s3 := sort.Reverse(s2)
	sort.Sort(s3)

	fmt.Printf("s1=%v\n", s1)
	fmt.Printf("s2=%v\n", s2)
	fmt.Printf("s3=%v\n", s3)


}
//out: [5 20 50 88 100]
