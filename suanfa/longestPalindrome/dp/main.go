package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s1 := []string{"a", "b", "c"}
	key,_ := json.Marshal(s1)
	fmt.Println(s1)
	fmt.Printf("%T", string(key))


	s := "cbbd"
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen < 2 {
		return s
	}
	dp := make([][]bool, strLen)
	for i := 0; i < strLen; i++ {
		dp[i] = make([]bool, strLen)
		dp[i][i] = true
	}
	longestStart := 0
	maxLength := 1
	for end := 1; end < strLen; end++ {
		for start := 0; start < end; start++ {
			//不相等 == false
			if s[start] != s[end] {
				dp[start][end] = false
			} else {
				//3个以下 == true
				if end - start + 1< 3 {
					dp[start][end] = true
				} else {
					//状态转移方程
					dp[start][end] = dp[start + 1][end - 1]
				}
			}
			if dp[start][end] &&  end - start + 1 > maxLength {
				longestStart = start
				maxLength = end -start + 1
			}
		}
	}
	return s[longestStart : longestStart + maxLength]
}
