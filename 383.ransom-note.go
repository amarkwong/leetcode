/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	mag := make(map[string]int)
	ran := make(map[string]int)
	for _, char := range ransomNote {
		if _, ok := ran[string(char)]; ok {
			ran[string(char)]++
		} else {
			ran[string(char)] = 1
		}
	}
	for _, char := range magazine {
		if _, ok := mag[string(char)]; ok {
			mag[string(char)]++
		} else {
			mag[string(char)] = 1
		}
	}
	for k, _ := range ran {
		if ran[k] <= mag[k] {
			continue
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
