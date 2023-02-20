package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target)) // Output: 4

	// 示例 2
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 3
	fmt.Println(search(nums, target)) // Output: -1

	// 示例 3
	nums = []int{-1}
	target = 0
	fmt.Println(search(nums, target)) // Output: -1
}

func search(nums []int, target int) int {
	// 解法一：二分查找
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] >= nums[0] {
			if target < nums[middle] && target >= nums[0] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if target > nums[middle] && target <= nums[len(nums)-1] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return -1
}
