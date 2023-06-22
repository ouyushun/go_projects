package main

import "fmt"

func main() {
	s := "ac"
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	maxLength := 0
	start := 0
	for i := 0; i < length - 1; i++ {
		len1 := judgeHuiwen(s, i, i)
		len2 := judgeHuiwen(s, i, i + 1)
		if max(len1, len2) > maxLength {
			maxLength = max(len1, len2)
			start = i - (maxLength - 1) / 2
		}
	}

	return s[start : start + maxLength]
}


func judgeHuiwen(s string, start, end int) int{
	for start <= end && start >= 0 && end < len(s) {
		if s[start] != s[end] {
			break
		}
		start--
		end++
	}

	return end - start  + 1 - 2
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}