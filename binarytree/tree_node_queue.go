package binarytree

// Queue ...
type TreeNodeQueue struct {
	data []*TreeNode
}

// NewQueue return new Queue
func NewTreeNodeQueue() *TreeNodeQueue {
	return &TreeNodeQueue{
		data: []*TreeNode{},
	}
}

// IsEmpty ...
func (q *TreeNodeQueue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek ...
func (q *TreeNodeQueue) Peek() *TreeNode {
	return q.data[len(q.data)-1]
}

// Enqueue ...
func (q *TreeNodeQueue) Enqueue(node *TreeNode) {
	q.data = append(q.data, node)
}

// Dequeue ...
func (q *TreeNodeQueue) Dequeue() *TreeNode {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}
