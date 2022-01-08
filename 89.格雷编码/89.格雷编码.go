/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
func grayCode(n int) []int {
	ans := make([]int, 1<<n)
	for i := 0; i < (1<<n); i++ {
		ans[i] = (i >> 1) ^ i
	}
	return ans
}
// @lc code=end

