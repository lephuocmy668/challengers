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

	newNode.Prev = lk.Tail
	lk.Tail.Next = newNode
	lk.Tail = newNode
}

// Prepend ...
func (lk *LinkedList) Prepend(newNode *Node) {
	if lk.Head != nil {
		lk.Head.Prev = newNode
	}
	newNode.Next = lk.Head
	lk.Head = newNode

	if lk.Tail == nil {
		lk.Tail = newNode
	}
}

// ToResult ...
func (lk *LinkedList) ToResult(k int) []int {
	result := []int{}
	temp := 0

	if lk.Tail == nil {
		return result
	}
	node := lk.Tail

	for node != nil {
		for _, e := range node.Elements {
			result = append(result, e.k)
			temp++
			if temp == k {
				return result
			}
		}
		node = node.Prev
	}
	return result
}
