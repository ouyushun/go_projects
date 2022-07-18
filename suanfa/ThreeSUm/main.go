package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "and"
	for _, item := range s {
		fmt.Printf("%T", item)
		fmt.Printf("%v", item)
	}


	nums := []int{-4, -1,-1,0,1,2}
	res := ThreeSum(nums)
	fmt.Println(res)
}


func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}


func ThreeSum(nums []int) [][]int  {

	sort.Ints(nums)
	var res [][]int
	l := 0
	r := len(nums) -1
	i := 1
	sum := 0
	for {
		fmt.Printf("l=%v ,i=%v, r=%v \n", l, i, r)
		if l == r  {
			break
		}
		if i > len(nums) - 1 {
			break
		}

		sum = nums[l] + nums[i] + nums[r]

		if sum == 0 {
			res = append(res, []int{nums[l], nums[i], nums[r]})
		}

		if sum > 0 {
			for ii := i+1; ii <= r; ii++ {
				if ii >= r{
					break
				}
				sum = nums[l] + nums[ii] + nums[r]
				if sum == 0 {
					res = append(res, []int{nums[l], nums[ii], nums[r]})
					break
				}
			}
			r = r - 1
		}

		if sum < 0 {
			for iii := i+1; iii <= r; iii++ {
				if iii >= r{
					break
				}
				sum = nums[l] + nums[iii] + nums[r]
				if sum == 0 {
					res = append(res, []int{nums[l], nums[iii], nums[r]})
					break
				}
			}
			l++
		}
		i++
	}

	return res
}