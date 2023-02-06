package main

import (
	"fmt"
)

func main() {
	// 示例 1
	s := "aa"
	p := "a"
	fmt.Println(isMatch(s, p)) // Output: false

	// 示例 2
	s = "aa"
	p = "a*"
	fmt.Println(isMatch(s, p)) // Output: true

	// 示例 3
	s = "ab"
	p = ".*"
	fmt.Println(isMatch(s, p)) // Output: true
}

func isMatch(s string, p string) bool {
	// 动态规划
	m, n := len(s)+1, len(p)+1
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
}
