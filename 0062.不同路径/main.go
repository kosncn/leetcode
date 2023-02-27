package main

import "fmt"

func main() {
	// 示例 1
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n)) // Output: 28

	// 示例 2
	m = 3
	n = 2
	fmt.Println(uniquePaths(m, n)) // Output: 3

	// 示例 3
	m = 7
	n = 3
	fmt.Println(uniquePaths(m, n)) // Output: 28

	// 示例 4
	m = 3
	n = 3
	fmt.Println(uniquePaths(m, n)) // Output: 6
}

func uniquePaths(m int, n int) int {
	// 解法一：动态规划
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
