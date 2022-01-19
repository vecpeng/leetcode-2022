/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
    var m = make(map[int]int)
    for i, num := range nums {
        v, ok := m[num]
        if ok {
            if (v - i) <= k && (i - v) <= k{
                return true
            } else {
                m[num] = i
            }
        } else {
            m[num] = i
        }
    }
    return false
}
// @lc code=end

