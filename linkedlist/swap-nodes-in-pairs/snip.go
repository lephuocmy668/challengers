package snip

import (
	"github.com/lephuocmy668/challengers/linkedlist"
)

// https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs(head *linkedlist.ListNode) *linkedlist.ListNode {
	first, last, next := doPair(head)

	resultNode := first
	node := next

	for node != nil {
		first, newLast, next := doPair(node)
		last.Next = first
		last = newLast
		node = next
	}
	return resultNode
}

func verify(head *linkedlist.ListNode) []int {
	result := swapPairs(head)
	return result.ToSlice()
}

func doPair(node *linkedlist.ListNode) (first *linkedlist.ListNode, last *linkedlist.ListNode, next *linkedlist.ListNode) {
	if node == nil {
		return nil, nil, nil
	}

	if node.Next == nil {
		return node, nil, nil
	}

	first = node.Next
	next = first.Next
	first.Next = node
	last = node
	last.Next = next

	return
}
