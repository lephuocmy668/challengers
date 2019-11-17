package rnnfeof

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	var pre *ListNode
	var last *ListNode
	next := head
	node := head
	temp := 0

	for node != nil {
		last = node
		temp += 1

		if temp > n-1 {
			next = next.Next
		}

		if temp == n {
			if node.Next != nil {
				pre = head
			}
		}

		if temp > n {
			if node.Next != nil {
				pre = pre.Next
			}
		}

		node = node.Next
	}

	if temp == n && n == 1 {
		return nil
	}

	if pre == nil {
		if next == nil {
			return last
		}
		return next
	}

	if next == nil {
		pre.Next = nil
	} else {
		pre.Next = next
	}
	return head
}
