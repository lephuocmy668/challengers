package reverse

// https://leetcode.com/problems/reverse-linked-list-ii/submissions/
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2.1 MB, less than 12.50% of Go online submissions for Reverse Linked List II.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	index := 1
	var preTail, newTail, newHead, preHead *ListNode

	originHead := head
	var prevNode *ListNode
	currentNode := head

	for currentNode != nil {
		node := currentNode
		currentNode = currentNode.Next

		if index == m {
			newHead = node
			preHead = prevNode
			newTail = node
			preTail = node.Next
		}
		if index == n {
			newHead = node
			preTail = node.Next
			node.Next = prevNode
			break
		}
		if index > m && index < n {
			newHead = node
			node.Next = prevNode
			preTail = node.Next
		}

		prevNode = node
		index++
	}

	if newHead != nil && preHead != nil {
		preHead.Next = newHead
	}
	if newTail != nil {
		newTail.Next = preTail
	}

	if m > 1 {
		return originHead
	}
	return newHead
}
