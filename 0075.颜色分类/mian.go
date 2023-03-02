package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums) // Output: [0, 0, 1, 1, 2, 2]

	// 示例 2
	nums = []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums) // Output: [0, 1, 2]
}

func sortColors(nums []int) {
	// 解法一：单指针
	//n := len(nums)
	//p := 0
	//for i := 0; i < n; i++ {
	//	if nums[i] == 0 {
	//		nums[i], nums[p] = nums[p], nums[i]
	//		p++
	//	}
	//}
	//for i := p; i < n; i++ {
	//	if nums[i] == 1 {
	//		nums[i], nums[p] = nums[p], nums[i]
	//		p++
	//	}
	//}

	// 解法二：双指针
	//n := len(nums)
	//p0, p1 := 0, 0
	//for i := 0; i < n; i++ {
	//	if nums[i] == 1 {
	//		nums[i], nums[p1] = nums[p1], nums[i]
	//		p1++
	//	} else if nums[i] == 0 {
	//		nums[i], nums[p0] = nums[p0], nums[i]
	//		if p0 < p1 {
	//			nums[i], nums[p1] = nums[p1], nums[i]
	//		}
	//		p0, p1 = p0+1, p1+1
	//	}
	//}

	// 解法三：双指针
	n := len(nums)
	p0, p2 := 0, n-1
	for i := 0; i <= p2; i++ {
		for i <= p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
