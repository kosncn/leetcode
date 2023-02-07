package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSwm(nums, target)) // Output: [0, 1]

	// 示例 2
	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSwm(nums, target)) // Output: [1, 2]

	// 示例 3
	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSwm(nums, target)) // Output: [0, 1]
}

func twoSwm(nums []int, target int) []int {
	// 解法一: 暴力
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i]+nums[j] == target {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//return []int{}

	// 解法二: 哈希
	number := make(map[int]int)
	for i, v := range nums {
		if index, ok := number[target-v]; ok {
			return []int{index, i}
		}
		number[v] = i
	}
	return []int{}
}
