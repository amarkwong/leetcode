/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */
// @lc code=start
func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] = -i
		} else {
			m[s[i]] = i
		}
	}
	minv := -1
	for _, v := range m {
		if v >= 0 && minv == -1 {
			minv = v
		} else if v >= 0 && v < minv {
			minv = v
		}
	}
	return minv
}

// @lc code=end
