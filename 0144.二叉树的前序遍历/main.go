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
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	fmt.Println(preorderTraversal(root)) // Output: [1, 2, 3]

	// 示例 2
	root = nil
	fmt.Println(preorderTraversal(root)) // Output: []

	// 示例 3
	root = &TreeNode{Val: 1}
	fmt.Println(preorderTraversal(root)) // Output: [1]

	// 示例 4
	root = &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	fmt.Println(preorderTraversal(root)) // Output: [1, 2]

	// 示例 5
	root = &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(preorderTraversal(root)) // Output: [1, 2]
}

func preorderTraversal(root *TreeNode) []int {
	// 解法一：回溯算法思路
	//var result []int
	//var traverse func(*TreeNode)
	//traverse = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	result = append(result, node.Val)
	//	traverse(node.Left)
	//	traverse(node.Right)
	//}
	//traverse(root)
	//return result

	// 解法二：动态规划思路
	var result []int
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}
