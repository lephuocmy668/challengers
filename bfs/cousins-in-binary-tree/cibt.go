package cibt

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Cousins in Binary Tree.
// TODO: romove duplicate code
// https://leetcode.com/problems/cousins-in-binary-tree/submissions/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
        return false
	}

    var nodeX *QueueNode
    var nodeY *QueueNode
    
	q := NewQueue()
	if root.Left != nil {
		q.Enqueue(&QueueNode{
			TreeNode: root.Left,
			Height:   1,
		})   
	} else {
		q.Enqueue(nil)  
	}
	
	if root.Right != nil {
		q.Enqueue(&QueueNode{
			TreeNode: root.Right,
			Height:   1,
		})   
	} else {
		q.Enqueue(nil)  
	}

	for !q.IsEmpty() {
		node1 := q.Dequeue()
		node2 := q.Dequeue()

		if node1 != nil && node1.TreeNode.Val == x && node2 != nil && node2.TreeNode.Val == y {
			return false
		}

		if node1 != nil && node1.TreeNode.Val == y && node2 != nil && node2.TreeNode.Val == x {
			return false
		}

		if node1 != nil {
            if node1.TreeNode.Val == x {
                nodeX = node1
            }
            
            if node1.TreeNode.Val == y {
                nodeY = node1
            }
            
            if node1.TreeNode.Left != nil {
                q.Enqueue(&QueueNode{
                    TreeNode: node1.TreeNode.Left,
                    Height:   node1.Height + 1,
                })   
            } else {
                q.Enqueue(nil)  
            }
            
            if node1.TreeNode.Right != nil {
                q.Enqueue(&QueueNode{
                    TreeNode: node1.TreeNode.Right,
                    Height:   node1.Height + 1,
                })   
            } else {
                q.Enqueue(nil)  
            }
		}

		if node2 != nil {
            if node2.TreeNode.Val == x {
                nodeX = node2
            }
            
            if node2.TreeNode.Val == y {
                nodeY = node2
            }
            
            if node2.TreeNode.Left != nil {
                q.Enqueue(&QueueNode{
                    TreeNode: node2.TreeNode.Left,
                    Height:   node2.Height + 1,
                })   
            } else {
                q.Enqueue(nil)  
            }
            
            if node2.TreeNode.Right != nil {
                q.Enqueue(&QueueNode{
                    TreeNode: node2.TreeNode.Right,
                    Height:   node2.Height + 1,
                })   
            } else {
                q.Enqueue(nil)  
            }
		}
	}
    
    if nodeX == nil || nodeY == nil {
        return false
    }
    
    if nodeX.Height != nodeY.Height {
        return false
    }
    
	return true
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