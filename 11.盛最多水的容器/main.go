package main

import "fmt"

func main() {
	// 示例 1
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) // Output: 49

	// 示例 2
	height = []int{1, 1}
	fmt.Println(maxArea(height)) // Output: 1
}

func maxArea(height []int) int {
	// 方法一: 暴力 (超时)
	//result := 0
	//length := len(height)
	//for i := 0; i < length; i++ {
	//	for j := i + 1; j < length; j++ {
	//		area := (j - i) * min(height[i], height[j])
	//		result = max(result, area)
	//	}
	//}
	//return result

	// 方法二: 双指针
	result := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		result = max(result, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
