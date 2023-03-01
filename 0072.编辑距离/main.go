package main

import "fmt"

//var memo [][]int

func main() {
	// 示例 1
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2)) // Output: 3

	// 示例 2
	word1 = "intention"
	word2 = "execution"
	fmt.Println(minDistance(word1, word2)) // Output: 5
}

func minDistance(word1 string, word2 string) int {
	// 解法一：动态规划（自顶向下）
	//m, n := len(word1), len(word2)
	//memo = make([][]int, m)
	//for i := 0; i < m; i++ {
	//	memo[i] = make([]int, n)
	//}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		memo[i][j] = -1
	//	}
	//}
	//return dp(word1, m-1, word2, n-1)

	// 解法二：动态规划（自底向上）
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i][j-1],   // 插入
					dp[i-1][j],   // 删除
					dp[i-1][j-1], // 替换
				) + 1
			}
		}
	}
	return dp[m][n]
}

//func dp(s1 string, i int, s2 string, j int) int {
//	if i <= -1 {
//		return j + 1
//	}
//	if j <= -1 {
//		return i + 1
//	}
//	if memo[i][j] != -1 {
//		return memo[i][j]
//	}
//	if s1[i] == s2[j] {
//		memo[i][j] = dp(s1, i-1, s2, j-1) // 跳过
//	} else {
//		memo[i][j] = min(
//			dp(s1, i, s2, j-1),   // 插入
//			dp(s1, i-1, s2, j),   // 删除
//			dp(s1, i-1, s2, j-1), // 替换
//		) + 1
//	}
//	return memo[i][j]
//}

func min(value ...int) int {
	minValue := value[0]
	for _, v := range value {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}
