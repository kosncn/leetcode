package main

import "fmt"

func main() {
	// 示例 1
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix)) // Output: 6

	// 示例 2
	matrix = [][]byte{}
	fmt.Println(maximalRectangle(matrix)) // Output: 0

	// 示例 3
	matrix = [][]byte{{'0'}}
	fmt.Println(maximalRectangle(matrix)) // Output: 0

	// 示例 4
	matrix = [][]byte{{'1'}}
	fmt.Println(maximalRectangle(matrix)) // Output: 1

	// 示例 5
	matrix = [][]byte{{'0', '0'}}
	fmt.Println(maximalRectangle(matrix)) // Output: 0
}

func maximalRectangle(matrix [][]byte) int {
	// 解法一：使用柱状图优化的暴力法
	//area := 0
	//if len(matrix) == 0 {
	//	return area
	//}
	//m, n := len(matrix), len(matrix[0])
	//left := make([][]int, m)
	//for i := 0; i < m; i++ {
	//	left[i] = make([]int, n)
	//	left[i][0] = int(matrix[i][0] - '0')
	//}
	//for i := 0; i < m; i++ {
	//	for j := 1; j < n; j++ {
	//		if matrix[i][j] == '1' {
	//			left[i][j] = left[i][j-1] + 1
	//		}
	//	}
	//}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if matrix[i][j] == '0' {
	//			continue
	//		}
	//		width := left[i][j]
	//		temp := left[i][j]
	//		for k := i - 1; k >= 0; k-- {
	//			width = min(width, left[k][j])
	//			temp = max(temp, (i-k+1)*width)
	//		}
	//		area = max(area, temp)
	//	}
	//}
	//return area

	// 解法二：单调栈
	area := 0
	if len(matrix) == 0 {
		return area
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
		left[i][0] = int(matrix[i][0] - '0')
	}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ {
		var stack []int
		up, down := make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			for len(stack) > 0 && left[stack[len(stack)-1]][j] >= left[i][j] {
				stack = stack[:len(stack)-1]
			}
			up[i] = -1
			if len(stack) > 0 {
				up[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
		stack = []int{}
		for i := m - 1; i >= 0; i-- {
			for len(stack) > 0 && left[stack[len(stack)-1]][j] >= left[i][j] {
				stack = stack[:len(stack)-1]
			}
			down[i] = m
			if len(stack) > 0 {
				down[i] = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
		for i := 0; i < m; i++ {
			area = max(area, (down[i]-up[i]-1)*left[i][j])
		}
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
