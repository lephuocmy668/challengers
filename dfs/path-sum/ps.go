package ps

import (
	"github.com/lephuocmy668/challengers/binarytree"
)

// https://leetcode.com/problems/path-sum/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *binarytree.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}

	sums := make(map[*binarytree.TreeNode]int)
	sums[root] = root.Val

	s := binarytree.NewTreeNodeStack()

l:
	for {
		for root != nil {
			tempNode := root
			s.Push(root)
			root = root.Right
			if lastSum, ok := sums[tempNode]; ok {
				if root != nil {
					tempSum := lastSum + root.Val
					sums[root] = tempSum
					if root.Left == nil && root.Right == nil && tempSum == sum {
						return true
					}
				}
			}
		}
		if s.IsEmpty() {
			break l
		}

		root = s.Pop()
		tempNode := root
		root = root.Left
		if lastSum, ok := sums[tempNode]; ok {
			if root != nil {
				tempSum := lastSum + root.Val
				sums[root] = tempSum
				if root.Left == nil && root.Right == nil && tempSum == sum {
					return true
				}
			}
		}
	}

	return false
}
