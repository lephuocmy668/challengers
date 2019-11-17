package atn

// https://leetcode.com/problems/add-two-numbers/submissions/
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
/*
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var startNode *ListNode
	var pre *ListNode
	node1 := l1
	node2 := l2
	temp := 0

	for node1 != nil {
		newNode := ListNode{
			Val: node1.Val + temp,
		}
		temp = 0

		if pre == nil {
			startNode = &newNode
		} else {
			pre.Next = &newNode
		}

		if node2 != nil {
			newNode.Val += node2.Val
			node2 = node2.Next
		}

		if newNode.Val > 9 {
			temp = newNode.Val / 10
			newNode.Val = newNode.Val % 10
		}

		node1 = node1.Next
		pre = &newNode
	}

	if node2 != nil {
		for node2 != nil {
			newNode := ListNode{
				Val: node2.Val + temp,
			}
			temp = 0

			if newNode.Val > 9 {
				temp = newNode.Val / 10
				newNode.Val = newNode.Val % 10
			}

			pre.Next = &newNode
			node2 = node2.Next
			pre = &newNode
		}
	}

	if temp > 0 {
		pre.Next = &ListNode{
			Val: temp,
		}
	}

	return startNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
