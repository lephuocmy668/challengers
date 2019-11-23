package st

//  https://leetcode.com/problems/symmetric-tree/
/*
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
//  https://leetcode.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmentric(root.Left, root.Right)
}

func checkSymmentric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return checkSymmentric(left.Left, right.Right) && checkSymmentric(left.Right, right.Left)
}
