/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 *
 * https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (38.97%)
 * Likes:    1922
 * Dislikes: 147
 * Total Accepted:    155.3K
 * Total Submissions: 398.5K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * Given a string s and a non-empty string p, find all the start indices of p's
 * anagrams in s.
 *
 * Strings consists of lowercase English letters only and the length of both
 * strings s and p will not be larger than 20,100.
 *
 * The order of output does not matter.
 *
 * Example 1:
 *
 * Input:
 * s: "cbaebabacd" p: "abc"
 *
 * Output:
 * [0, 6]
 *
 * Explanation:
 * The substring with start index = 0 is "cba", which is an anagram of "abc".
 * The substring with start index = 6 is "bac", which is an anagram of
 * "abc".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * s: "abab" p: "ab"
 *
 * Output:
 * [0, 1, 2]
 *
 * Explanation:
 * The substring with start index = 0 is "ab", which is an anagram of "ab".
 * The substring with start index = 1 is "ba", which is an anagram of "ab".
 * The substring with start index = 2 is "ab", which is an anagram of "ab".
 *
 *
 */
// @lc code=start
import "reflect"

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) == 0 || len(s) < len(p) {
		return nil
	}

	mp, l := make(map[string]int), len(p)
	ms := make(map[string]int)
	res := []int{}
	for _, char := range p {
		mp[string(char)]++
	}
	for _, char := range s[:l] {
		ms[string(char)]++
	}
	for i := 0; i < len(s)-l+1; i++ {
		if reflect.DeepEqual(ms, mp) {
			res = append(res, i)
		}
		ms[string(s[i])]--
		if ms[string(s[i])] == 0 {
			delete(ms, string(s[i]))
		}
		if i+l < len(s) {
			ms[string(s[i+l])]++
		}
	}
	return res
}

// @lc code=end
