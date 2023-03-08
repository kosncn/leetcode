package main

import "fmt"

func main() {
	// 示例 1
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // Output: true

	// 示例 2
	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "SEE"
	fmt.Println(exist(board, word)) // Output: true

	// 示例 3
	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCB"
	fmt.Println(exist(board, word)) // Output: false

	// 示例 4
	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCESEEEFS"
	fmt.Println(exist(board, word)) // Output: true
}

func exist(board [][]byte, word string) bool {
	// 解法一：深度优先遍历
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for k := 0; k < m; k++ {
		visited[k] = make([]bool, n)
	}
	var dfs func(int, int, int) bool
	dfs = func(x int, y int, index int) bool {
		if index == len(word) {
			return true
		}
		if x < 0 || x > m-1 || y < 0 || y > n-1 || visited[x][y] {
			return false
		}
		if board[x][y] != word[index] {
			return false
		}
		visited[x][y] = true
		result := dfs(x-1, y, index+1) || dfs(x+1, y, index+1) || dfs(x, y-1, index+1) || dfs(x, y+1, index+1)
		visited[x][y] = false
		return result
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
