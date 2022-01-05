/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode)  {
	var prev, x, y, pred *TreeNode
	for root != nil {
		if root.Left != nil {
			prev = root.Left
			for prev.Right != nil && prev.Right != root {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = root
				root = root.Left
			} else {
				if pred != nil && root.Val < pred.Val {
                    y = root
                    if x == nil {
                        x = pred
                    }
                }
				pred = root
				root = root.Right
				prev.Right = nil
			}
		} else {
			if pred != nil && root.Val < pred.Val {
                y = root
                if x == nil {
                    x = pred
                }
            }
            pred = root
            root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}
// @lc code=end

