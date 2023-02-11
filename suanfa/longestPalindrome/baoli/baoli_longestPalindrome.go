package main

import "fmt"

func main() {
	res := longestPalindrome("aabdbdaa")
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen <= 1 {
		return s
	}
	maxLen := 0
	start := 0
	for i := 0; i < strLen; i++ {
		for j := i; j < strLen; j++ {
			if huiwen(s, i, j) {
				length := j - i + 1
				if length > maxLen {
					maxLen = length
					start = i
				}
			}
		}
	}

	return s[start : start + maxLen]
}


func huiwen(s string, left, right int) bool {
	//判断回文， 两个指针从两边往中间走
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
