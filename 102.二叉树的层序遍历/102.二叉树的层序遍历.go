/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		for _, tree := range queue {
			level = append(level, tree.Val)
			if tree.Left != nil {
				queue = append(queue, tree.Left)
			}
			if tree.Right != nil {
				queue = append(queue, tree.Right)
			}
			queue = queue[1:]
		}
		res = append(res, level)
	}
	return res
}
// @lc code=end

