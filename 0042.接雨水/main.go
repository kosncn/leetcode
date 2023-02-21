package main

import "fmt"

func main() {
	// 示例 1
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height)) //Output: 6

	// 示例 2
	height = []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height)) //Output: 9
}

func trap(height []int) int {
	// 解法一：双指针
	result := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] <= height[right] {
			leftMax = max(leftMax, height[left])
			result += leftMax - height[left]
			left++
		} else {
			rightMax = max(rightMax, height[right])
			result += rightMax - height[right]
			right--
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
