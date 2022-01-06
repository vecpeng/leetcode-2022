/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = postorder[len(postorder)-1]
	for i, item := range inorder {
		if item == root.Val {
			root.Left = buildTree(inorder[0:i], postorder[0:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
			break
		}
	}
	return root
}
// @lc code=end

