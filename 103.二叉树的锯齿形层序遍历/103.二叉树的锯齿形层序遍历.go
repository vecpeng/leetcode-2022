/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	rev := false
	for len(queue) > 0 {
		level := []int{}
		for _, tree := range queue {
			if !rev {
				level = append(level, tree.Val)
			} else {
				level = append([]int{tree.Val}, level...)
			}
			if tree.Left != nil {
				queue = append(queue, tree.Left)
			}
			if tree.Right != nil {
				queue = append(queue, tree.Right)
			}
			queue = queue[1:]
		}
		rev = !rev
		res = append(res, level)
	}
	return res
}
// @lc code=end

