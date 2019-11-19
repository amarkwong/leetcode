/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

import "strings"

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	//recursive
	//find the first match of the prefix in wordDict
	//if doesn't exist, return false
	//else remove the prefix
	//repeat the same move till the string is empty
	tmps := s
	for count := 0; len(tmps) > 0; {
		for _, word := range wordDict {
			if strings.HasPrefix(tmps, word) {
				strings.TrimPrefix(tmps, word)
				count++
			}
		}
		if count == 0 {
			return false
		}
		count = 0
	}
	return true
}

// @lc code=end
