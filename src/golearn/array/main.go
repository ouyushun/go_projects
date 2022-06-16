package main

import "fmt"

func main() {

	//定义数组的几种方法
	var arr1 [5]int
	fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	var grid [2][2][2][2]bool
	fmt.Println(grid)

	for _, v := range arr3 {
		fmt.Println(v)
	}

	//数组是值类型
	PrintArray(&[3]int{1, 2, 3})

}

func PrintArray(arr *[3]int) {
	fmt.Println("--------")
	arr[0] = 123
	fmt.Println(arr)
}