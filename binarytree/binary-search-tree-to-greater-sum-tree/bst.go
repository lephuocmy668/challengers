package bst

import "github.com/lephuocmy668/challengers/binarytree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
func bstToGst(root *binarytree.TreeNode) *binarytree.TreeNode {
	sum := 0
	res := root

	s := binarytree.NewTreeNodeStack()
l:
	for {
		for root != nil {
			s.Push(root)
			root = root.Right
		}
		if s.IsEmpty() {
			break l
		}

		root = s.Pop()
		sum += root.Val
		root.Val = sum
		root = root.Left
	}

	return res
}
