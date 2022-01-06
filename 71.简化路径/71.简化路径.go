/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	stack := []string{}
    strs := strings.Split(path, "/")
    for _, str := range strs {
        if len(stack) > 0 && str == ".." {
            stack = stack[:len(stack)-1]
        } else if str != "." && str != "" && str != ".."{
            stack = append(stack, str)
        }
    }

    res := ""
	for _, s := range stack {
        res += "/" + s
    }
    if res == "" {
        return "/"
    }
	return res
}
// @lc code=end

