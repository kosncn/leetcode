package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [1, 3, 2]

	// 示例 2
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [1, 2, 3]

	// 示例 3
	nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [1, 5, 1]
}

func nextPermutation(nums []int) {
	// 解法一：两遍扫描
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= i+1 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for left, right := i+1, n-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}
