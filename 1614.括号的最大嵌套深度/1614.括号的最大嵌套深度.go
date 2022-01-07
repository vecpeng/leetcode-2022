/*
 * @lc app=leetcode.cn id=1614 lang=golang
 *
 * [1614] 括号的最大嵌套深度
 */

// @lc code=start
func maxDepth(s string) int {
	size := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			size++
		} else if s[i] == ')' {
			size--
		}
		if size > ans {
			ans = size
		}
	}
	return ans
}
// @lc code=end

