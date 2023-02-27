package main

import (
	"fmt"
	"sort"
)

func main() {
	// 示例 1
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // Output: [[1, 6], [8, 10], [15, 18]]

	// 示例 2
	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals)) // Output: [[1, 5]]

	// 示例 3
	intervals = [][]int{{1, 4}, {0, 4}}
	fmt.Println(merge(intervals)) // Output: [[0, 4]]

	// 示例 3
	intervals = [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	fmt.Println(merge(intervals)) // Output: [[0, 4]]
}

func merge(intervals [][]int) [][]int {
	// 解法一：排序
	var result [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		index := len(result) - 1
		if i > 0 && intervals[i][0] <= result[index][1] {
			result[index][1] = max(result[index][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
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
