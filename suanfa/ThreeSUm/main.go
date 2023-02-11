package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -1,-1,0,1,2}
	res := threeSum(nums)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if (len(nums) == 0) {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums) - 1; i++ {
		if (nums[i] > 0) {
			return res
		}
		//去重复的i
		if (i > 0 && nums[i] == nums[i - 1]) {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for (l < r) {
			sum := nums[i] + nums[l] + nums[r]
			if (sum > 0) {
				r--
			} else if (sum < 0) {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				//去重
				for (l < r && nums[l] == nums[l-1]) {
					l++
				}
				for (l < r && nums[r] == nums[r+1]) {
					r--
				}
			}
		}
	}
	return res
}