package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i int, j int) bool { return pq[i].Val < pq[j].Val }

func (pq PriorityQueue) Swap(i int, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return node
}

func main() {
	// 示例 1
	lists := []*ListNode{ // lists: [[1, 4, 5], [1, 3, 4], [2, 6]]
		{
			Val: 1,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 5},
			},
		},
		{
			Val: 1,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 4},
			},
		},
		{
			Val:  2,
			Next: &ListNode{Val: 6},
		},
	}
	fmt.Println(printListNode(mergeKLists(lists))) // Output: [1, 1, 2, 3, 4, 4, 5, 6]

	// 示例 2
	lists = []*ListNode{}                          // lists: []
	fmt.Println(printListNode(mergeKLists(lists))) // Output: []

	// 示例 3
	lists = []*ListNode{nil}                       // lists: [[]]
	fmt.Println(printListNode(mergeKLists(lists))) // Output: []
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 解法一：顺序合并
	//var result *ListNode
	//for i := 0; i < len(lists); i++ {
	//	result = mergeTowLists(result, lists[i])
	//}
	//return result

	// 解法二：优先级队列
	if len(lists) <= 0 {
		return nil
	}
	pq := new(PriorityQueue)
	dummy := new(ListNode)
	p := dummy
	for _, v := range lists {
		if v != nil {
			heap.Push(pq, v)
		}
	}
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		p.Next = node
		p = p.Next
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}
	return dummy.Next
}

func mergeTowLists(list1, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p, p1, p2 := dummy, list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}

func printListNode(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
