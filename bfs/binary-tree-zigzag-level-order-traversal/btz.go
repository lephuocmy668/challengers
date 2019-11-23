package btz

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/submissions/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	height := 1
	st := newStack()
	st.push(root)

	for !st.isEmpty() {
		arr := []int{}
		newSt := newStack()

		for !st.isEmpty() {
			node := st.pop()
			if node != nil {
				arr = append(arr, node.Val)
				if height%2 == 0 {
					newSt.push(node.Right)
					newSt.push(node.Left)
				} else {
					newSt.push(node.Left)
					newSt.push(node.Right)
				}
			}
		}

		if len(arr) > 0 {
			res = append(res, arr)
		}

		st = newSt
		height++
	}

	return res
}

type stack struct {
	values []*TreeNode
}

func newStack() *stack {
	return &stack{
		values: []*TreeNode{},
	}
}

func (s *stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *stack) push(e *TreeNode) {
	s.values = append(s.values, e)
}

func (s *stack) pop() *TreeNode {
	result := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return result
}

func (s *stack) peek() *TreeNode {
	return s.values[len(s.values)-1]
}
