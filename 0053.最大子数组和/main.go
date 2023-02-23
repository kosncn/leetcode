package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums)) // Output: 6

	// 示例 2
	nums = []int{1}
	fmt.Println(maxSubArray(nums)) // Output: 1

	// 示例 3
	nums = []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums)) // Output: 23

	// 示例 3
	nums = []int{-1}
	fmt.Println(maxSubArray(nums)) // Output: -1
}

func maxSubArray(nums []int) int {
	// 解法一：动态规划
	//n := len(nums)
	//dp := make([]int, n)
	//dp[0] = nums[0]
	//for i := 1; i < n; i++ {
	//	dp[i] = max(nums[i], nums[i]+dp[i-1])
	//}
	//result := dp[0]
	//for i := 0; i < n; i++ {
	//	result = max(result, dp[i])
	//}
	//return result

	// 解法二：动态规划（压缩空间）
	n := len(nums)
	dp0, dp1 := nums[0], 0
	result := nums[0]
	for i := 1; i < n; i++ {
		dp1 = max(nums[i], nums[i]+dp0)
		result = max(result, dp1)
		dp0 = dp1
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
