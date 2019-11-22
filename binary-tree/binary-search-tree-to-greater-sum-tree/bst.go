package bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	res := root

	s := newStack()
l:
	for {
		for root != nil {
			s.push(root)
			root = root.Right
		}
		if s.isEmpty() {
			break l
		}

		root = s.pop()
		sum += root.Val
		root.Val = sum
		root = root.Left
	}

	return res
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	nodes []*TreeNode
}

func newStack() *stack {
	return &stack{
		nodes: []*TreeNode{},
	}
}

func (s *stack) isEmpty() bool {
	return len(s.nodes) == 0
}

func (s *stack) push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
}

func (s *stack) pop() *TreeNode {
	result := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return result
}

func (s *stack) peek() *TreeNode {
	return s.nodes[len(s.nodes)-1]
}
