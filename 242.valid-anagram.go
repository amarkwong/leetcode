import "reflect"

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		ms, mt := make(map[string]int), make(map[string]int)
		for _, v := range s {
			ms[string(v)]++
		}
		for _, v := range t {
			mt[string(v)]++
		}
		if reflect.DeepEqual(ms, mt) {
			return true
		} else {
			return false
		}
	}
}

// @lc code=end
