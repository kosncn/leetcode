package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 示例 1
	l1 := &ListNode{ // l1: [2, 4, 3]
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{Val: 3},
		},
	}
	l2 := &ListNode{ // l2: [5, 6, 4]
		Val: 5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{Val: 4},
		},
	}
	fmt.Println(printListNode(addTwoNumbers(l1, l2))) // Output: [7, 0, 8]

	// 示例 2
	l1 = &ListNode{Val: 0}                            // l1: [0]
	l2 = &ListNode{Val: 0}                            // l2: [0]
	fmt.Println(printListNode(addTwoNumbers(l1, l2))) // Output: [0]

	// 示例 3
	l1 = &ListNode{ // l1: [9, 9, 9, 9, 9, 9, 9]
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{Val: 9},
						},
					},
				},
			},
		},
	}
	l2 = &ListNode{ // l2: [9, 9, 9, 9]
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: &ListNode{Val: 9},
			},
		},
	}

	fmt.Println(printListNode(addTwoNumbers(l1, l2))) // Output: [8, 9, 9, 9, 0, 0, 0, 1]

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 解法一: 双指针
	result := new(ListNode)
	p, p1, p2 := result, l1, l2
	bit := 0
	for p1 != nil || p2 != nil || bit != 0 {
		if p1 != nil {
			bit += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			bit += p2.Val
			p2 = p2.Next
		}
		p.Next = &ListNode{Val: bit % 10}
		p = p.Next
		bit /= 10
	}
	return result.Next
}

func printListNode(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
