package main

import "fmt"

func main() {
	// 示例 1
	n := 2
	fmt.Println(climbStairs(n)) // Output: 2

	// 示例 2
	n = 3
	fmt.Println(climbStairs(n)) // Output: 3

	// 示例 3
	n = 1
	fmt.Println(climbStairs(n)) // Output: 3
}

func climbStairs(n int) int {
	// 解法一：动态规划
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
