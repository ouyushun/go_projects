package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 1}
	res := greddy(arr)
	fmt.Println(res)
}

//贪心
func greddy(nums []int) int {
	step := 0
	start :=0
	end := 0
	for start < len(nums) - 1 {
		fmt.Println("start = ", start)
		step++
		//一步到终点
		if start + nums[start] >= len(nums) -1 {
			break
		}
		maxPos := 0
		for i := start + 1; i <= start + nums[start] && i <= len(nums) -1; i++ {
			if nums[i] + i >= maxPos {
				maxPos = max(maxPos, nums[i] + i)
				end = i
			}
		}
		start = end
	}
	return step
}

//超时
func dp(start int, nums []int) int{
	dict := make(map[int]int, len(nums))
	steps := 1 << 31 - 1
	if start >= len(nums)-1 {
		return 0
	}
	currVal := nums[start]
	for j := start +1;j <= start + currVal; j++ {
		if _, ok := dict[j] ; !ok{
			steps = min(dp(j, nums) +1, steps)
			dict[j] = steps
		}
		steps = dict[j]
	}
	return steps
}

/*
作者：zwjago
链接：https://leetcode.cn/problems/jump-game-ii/solution/tiao-yue-you-xi-2cong-chao-shi-dao-100co-zzfq/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func jump(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	dp[size-1] = 0
	for i := size-2; i >= 0; i-- {
		dp[i] = 1 << 31 - 1
		for x := 1; x <= nums[i]; x++ {
			if i + x < size {
				dp[i] = min(dp[i], 1+ dp[i+x])
			}
		}
	}
	return dp[0]
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
