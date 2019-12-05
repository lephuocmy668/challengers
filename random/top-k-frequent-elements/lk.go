package te

// LinkedList ...
type LinkedList struct {
	Head *Node
	Tail *Node
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
func (lk *LinkedList) Append(newNode *Node) {
	if lk.Head == nil {
		lk.Head = newNode
		lk.Tail = newNode
		return
	}

	lk.Tail.Next = newNode
	lk.Tail = newNode
}
