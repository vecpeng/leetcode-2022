/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	num := make([]int, n + 1)
	num[0] = 1
	num[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			num[i] +=  num[i - j] * num[j - 1]
		}
	}

	return num[n]
}
// @lc code=end

