package mdobt

// https://leetcode.com/problems/minimum-depth-of-binary-tree/submissions/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
	q := NewQueue()
    if root.Right == nil && root.Left == nil {
        return 1
    }
    
    if root.Right != nil {
        q.Enqueue(&QueueNode{
            TreeNode: root.Right,
            Height:   2,
        })
    }

    if root.Left != nil {
        q.Enqueue(&QueueNode{
            TreeNode: root.Left,
            Height:   2,
        })
    }

	for !q.IsEmpty() {
		node := q.Dequeue()
        if node.TreeNode.Left == nil && node.TreeNode.Right == nil {
            return node.Height
        }

		if node.TreeNode.Left != nil {
			q.Enqueue(&QueueNode{
				TreeNode: node.TreeNode.Left,
				Height:   node.Height + 1,
			})
        }

		if node.TreeNode.Right != nil {
			q.Enqueue(&QueueNode{
				TreeNode: node.TreeNode.Right,
				Height:   node.Height + 1,
			})
		} 
	}

	return 0
}

// QueueNode
type QueueNode struct {
	TreeNode *TreeNode
	Height   int
}

// Queue ...
type Queue struct {
	data []*QueueNode
}

// NewQueue return new Queue
func NewQueue() *Queue {
	return &Queue{
		data: []*QueueNode{},
	}
}

// IsEmpty ...
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek ...
func (q *Queue) Peek() *QueueNode {
	return q.data[len(q.data)-1]
}

// Enqueue ...
func (q *Queue) Enqueue(node *QueueNode) {
	q.data = append(q.data, node)
}

// Dequeue ...
func (q *Queue) Dequeue() *QueueNode {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

