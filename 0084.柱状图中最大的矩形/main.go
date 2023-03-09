package main

import "fmt"

func main() {
	// 示例 1
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights)) // Output: 10

	// 示例 2
	heights = []int{2, 4}
	fmt.Println(largestRectangleArea(heights)) // Output: 4

	// 示例 3
	heights = []int{0, 1, 0, 2, 0, 3, 0}
	fmt.Println(largestRectangleArea(heights)) // Output: 3

	// 示例 4
	heights = []int{3, 6, 5, 7, 4, 8, 1, 0}
	fmt.Println(largestRectangleArea(heights)) // Output: 20
}

func largestRectangleArea(heights []int) int {
	// 解法一：暴力（超时）
	//area := 0
	//n := len(heights)
	//for i := 0; i < n; i++ {
	//	hMin := heights[i]
	//	for j := i; j < n; j++ {
	//		hMin = min(hMin, heights[j])
	//		area = max(area, (j-i+1)*hMin)
	//	}
	//}
	//return area

	// 解法二：暴力（超时）
	//area := 0
	//n := len(heights)
	//for i := 0; i < n; i++ {
	//	left := i
	//	for left >= 0 && heights[left] >= heights[i] {
	//		left--
	//	}
	//	right := i
	//	for right < n && heights[right] >= heights[i] {
	//		right++
	//	}
	//	area = max(area, (right-left-1)*heights[i])
	//}
	//return area

	// 解法三：单调栈
	//var stack []int
	//n := len(heights)
	//left := make([]int, n)
	//for i := 0; i < n; i++ {
	//	for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
	//		stack = stack[:len(stack)-1]
	//	}
	//	if len(stack) == 0 {
	//		left[i] = -1
	//	} else {
	//		left[i] = stack[len(stack)-1]
	//	}
	//	stack = append(stack, i)
	//}
	//stack = []int{}
	//right := make([]int, n)
	//for i := n - 1; i >= 0; i-- {
	//	for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
	//		stack = stack[:len(stack)-1]
	//	}
	//	if len(stack) == 0 {
	//		right[i] = n
	//	} else {
	//		right[i] = stack[len(stack)-1]
	//	}
	//	stack = append(stack, i)
	//}
	//area := 0
	//for i := 0; i < n; i++ {
	//	area = max(area, (right[i]-left[i]-1)*heights[i])
	//}
	//return area

	// 解法四：单调栈
	var stack []int
	area := 0
	slice := append([]int{0}, heights...)
	slice = append(slice, 0)
	for i := 0; i < len(slice); i++ {
		for len(stack) > 0 && slice[stack[len(stack)-1]] > slice[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area = max(area, (i-stack[len(stack)-1]-1)*slice[index])
		}
		stack = append(stack, i)
	}
	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
