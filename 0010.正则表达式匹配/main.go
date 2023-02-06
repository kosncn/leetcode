package main

import "fmt"

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
	// 方法一: 动态规划
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for j := 1; j < n+1; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] != p[j-2] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}
