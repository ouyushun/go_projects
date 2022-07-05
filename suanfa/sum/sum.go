package main

import "fmt"

func main () {
	nums := []int{3,3}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}


func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[v] = i
	}
	fmt.Println(m)
	for index, val := range nums {
		answer := target - val
		if _, ok := m[answer]; ok {
			if index == m[answer] {
				continue
			}
			res = append(res, index)
		}
	}
	return res
}
