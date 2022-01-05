/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	return generate(n, 0);
}

func generate(n, add int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{nil}
	}
	trees := []*TreeNode{}
	for i := 0; i < n; i++ {
		for _, t := range generate(i, add) {
			for _, t1 := range generate(n - i - 1, i + 1 + add) {
				tree := &TreeNode{i + 1 + add, t, t1}
				trees = append(trees, tree)
			}
		}
	}
	return trees
}
// @lc code=end

