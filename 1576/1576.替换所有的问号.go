/*
 * @lc app=leetcode.cn id=1576 lang=golang
 *
 * [1576] 替换所有的问号
 */

// @lc code=start
func modifyString(s string) string {
	al := [26]int{}
	for _, ch := range s {
		if ch != '?' {
			al[ch - 'a'] = 1
		}
	}
	fmt.Printf("%v\n", s)
	for strings.Contains(s, "?") {
		for i, a := range al {
			if a == 0 {
				old := "?"
				new := string(i + 'a')
				s = strings.Replace(s, old, new, 1)
			}
		}
	}

	return s
}
// @lc code=end

