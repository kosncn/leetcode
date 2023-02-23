package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums)) // Output: true

	// 示例 2
	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums)) // Output: false
}

func canJump(nums []int) bool {
	// 解法一：贪心算法
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
