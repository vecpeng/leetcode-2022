/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelper(root.Left, root.Right)
}
func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	}
	if right == nil && left != nil {
		return false
	}
	if right == nil && left == nil {
		return true
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}
// @lc code=end

