package main

import "fmt"

func main() {
	slice := []int{3,3,4}
	countMap := make(map[int]int)
	k := 0
	for _, value := range slice {
		if countMap[value] <= 2 {
			slice[k] = value
			k++
		}
		countMap[value]++
	}
	slice = slice[:k]
	fmt.Println(slice) //[1 2 4 5 7 8]
}
