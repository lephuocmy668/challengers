package bsti

// Runtime: 60 ms, faster than 11.38% of Go online submissions for Binary Search Tree Iterator.
// Memory Usage: 11.6 MB, less than 75.00% of Go online submissions for Binary Search Tree Iterator.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	Stack *TreeNodeStack
}

func Constructor(root *TreeNode) BSTIterator {
	s := NewTreeNodeStack()

	node := root
	for node != nil {
		s.Push(node)
		node = node.Left
	}

	return BSTIterator{
		Stack: s,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if !this.HasNext() {
		return 0
	}

	node := this.Stack.Pop()
	result := node.Val

	node = node.Right
	for node != nil {
		this.Stack.Push(node)
		node = node.Left
	}

	return result
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return !this.Stack.IsEmpty()
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

type TreeNodeStack struct {
	nodes []*TreeNode
}

func NewTreeNodeStack() *TreeNodeStack {
	return &TreeNodeStack{
		nodes: []*TreeNode{},
	}
}

func (s *TreeNodeStack) IsEmpty() bool {
	return len(s.nodes) == 0
}

func (s *TreeNodeStack) Push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
}

func (s *TreeNodeStack) Pop() *TreeNode {
	result := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return result
}

func (s *TreeNodeStack) Peek() *TreeNode {
	return s.nodes[len(s.nodes)-1]
}
