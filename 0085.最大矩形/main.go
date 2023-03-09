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
	// 解法一：
	if len(matrix) == 0 {
		return 0
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
	area := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			width := left[i][j]
			temp := left[i][j]
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				temp = max(temp, (i-k+1)*width)
			}
			area = max(area, temp)
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
