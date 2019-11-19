/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	str := s
	max, curLen, dupIdx := 0, 0, 0
	for idx, char := range str {
		if _, ok := m[string(char)]; ok {
			dupIdx = m[string(char)]
			s = str[dupIdx+1:]
			for k := range m {
				if m[k] <= dupIdx {
					delete(m, k)
				}
			}
			m[string(char)] = idx
			curLen = idx - dupIdx
			continue
		} else {
			m[string(char)] = idx
			curLen++
			if curLen > max {
				max++
			}
		}
	}
	return max
}

// @lc code=end
