package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 示例 1
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(maxDepth(root)) // Output: 3
}

func maxDepth(root *TreeNode) int {
	// 解法一：回溯算法思路
	//var depth int
	//var layer int
	//var traverse func(*TreeNode)
	//traverse = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	layer++
	//	depth = max(depth, layer)
	//	traverse(node.Left)
	//	traverse(node.Right)
	//	layer--
	//}
	//traverse(root)
	//return depth

	// 解法二：动态规划思路
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return 1 + max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
