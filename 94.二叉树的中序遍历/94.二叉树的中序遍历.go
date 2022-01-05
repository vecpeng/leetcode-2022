/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for root != nil || len(stack) > 0 {
		// 遍历左子树，依次加入栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 最右边的节点
		root = stack[len(stack) - 1]
		fmt.Println(root.Val)
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}
// @lc code=end

