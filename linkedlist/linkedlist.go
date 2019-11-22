package linkedlist

// ListNode represents node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ToSlice ...
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

// LinkedList ...
type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

// NewLinkedList returns new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// IsEmpty ...
func (lk *LinkedList) IsEmpty() bool {
	return lk.Head == nil
}

// Append ...
func (lk *LinkedList) Append(val int) {
	newNode := &ListNode{
		Val: val,
	}

	if lk.Head == nil {
		lk.Head = newNode
		lk.Tail = newNode
		return
	}

	lk.Tail.Next = newNode
	lk.Tail = newNode
}

// Prepend ...
func (lk *LinkedList) Prepend(val int) {
	newNode := &ListNode{
		Val: val,
	}
	lk.Head = newNode

	if lk.Tail == nil {
		lk.Tail = newNode
	}
}

// DeleteTail ...
func (lk *LinkedList) DeleteTail() {
	if lk.Head == lk.Tail {
		lk.Head = nil
		lk.Tail = nil
		return
	}

	preTail := lk.Head
	for preTail.Next != nil {
		if preTail.Next.Next == nil {
			preTail.Next = nil
		} else {
			preTail = preTail.Next
		}
	}
	lk.Tail = preTail
}

// DeleteHead ...
func (lk *LinkedList) DeleteHead() {
	if lk.Head == nil {
		return
	}

	if lk.Head.Next == nil {
		lk.Head = nil
		lk.Tail = nil
	} else {
		lk.Head = lk.Head.Next
	}
}
