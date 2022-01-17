/*
 * @lc app=leetcode.cn id=1220 lang=golang
 *
 * [1220] 统计元音字母序列的数目
 */

// @lc code=start
func countVowelPermutation(n int) int {
    // dp := make([]int, 5)
    // for i := 0; i < n; i++ {
    //     dp[0][i] = dp[1][i-1] + dp[2][i-1] + dp[4][i-1]
    //     dp[1][i] = dp[0][i-1] + dp[2][i-1]
    //     dp[2][i] = dp[3][i-1] + dp[1][i-1]
    //     dp[3][i] = dp[2][i-1]
    //     dp[4][i] = dp[2][i-1] + dp[3][i-1]
    // }
    const mod int = 1e9 + 7
    a := 1
    e := 1
    ii := 1
    o := 1
    u := 1
    for i := 1; i < n; i++ {
        at := e + ii + u
        et := a + ii
        it := e + o
        ot := ii
        ut := ii + o
        a = at % mod
        e = et % mod
        ii = it % mod
        o = ot % mod
        u = ut
    }
    return (a + e + ii + o + u) % (mod)
}
// @lc code=end

