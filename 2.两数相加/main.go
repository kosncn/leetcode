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
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val) // Output: [7, 0, 8]
		result = result.Next
	}

	// 示例 2
	l1 = &ListNode{Val: 0} // l1: [0]
	l2 = &ListNode{Val: 0} // l2: [0]
	result = addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val) // Output: [0]
		result = result.Next
	}

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
	result = addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val) // Output: [8, 9, 9, 9, 0, 0, 0, 1]
		result = result.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	bit := 0
	for l1 != nil || l2 != nil || bit != 0 {
		if l1 != nil {
			bit += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			bit += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{Val: bit % 10}
		p = p.Next
		bit /= 10
	}
	return dummy.Next
}
