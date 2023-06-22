package main

func main() {

}


func maxArea(nums []int) int {
	s := 0
	l := 0
	r := 0

	for l < r {
		s = max(s, min(nums[l], nums[r]) * (r - l))
		if  nums[l] > nums[r] {
			r--
		} else {
			l++
		}
	}

	return s
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}