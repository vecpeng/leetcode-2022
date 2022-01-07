/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
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
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
	return res
}
// @lc code=end

