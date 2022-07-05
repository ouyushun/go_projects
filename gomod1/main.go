package main


import "fmt"

func main() {

	var a [10]int
	s := a[0:4]
	fmt.Printf("%T", s)

	s[1] = 4
	fmt.Println(a)


	var arr = [2]int{1,4}
	fmt.Printf("%T", arr)
	arr2 := arr
	arr2[1] = 6
	fmt.Println(arr)
	fmt.Println(arr2)

}
