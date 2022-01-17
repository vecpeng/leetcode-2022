/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	return dfs(num, 0, 0, 0, 0)
}

func dfs(num string, index int, count int, prevprev int, prev int) bool {
	if count >= len(num) {
		return count > 2
	}

	current := 0
	for i := index; i < len(num); i++ {
		n, _ := strconv.Atoi(string(num[i]))
		if num[index] == '0' && i > index {
			return false
		}
		current = current * 10 + n

		if count > 2 {
			if current > (prevprev + prev) {
				return false
			}
			if current < (prevprev + prev) {
				continue
			}
		}

		if dfs(num, i+1, count+1, prev, current) {
			return true
		}
	}
	return false
}
// @lc code=end

