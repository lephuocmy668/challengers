package linkedlist

// ListNode represents node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) ToSlice() []int {
	result := []int{}
	if head == nil {
		return result
	}
	node := head

	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
