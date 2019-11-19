/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 *
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (36.79%)
 * Likes:    518
 * Dislikes: 71
 * Total Accepted:    91.1K
 * Total Submissions: 247.7K
 * Testcase Example:  '5\n7'
 *
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND
 * of all numbers in this range, inclusive.
 *
 * Example 1:
 *
 *
 * Input: [5,7]
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: [0,1]
 * Output: 0
 */

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	if m == n {
		return m
	}
	start, end := 1, 2
	for index := 0; index < 32; index++ {
		start, end = start<<1, end<<1
		if m >= start && m < end {
			if n >= start && n < end {
				if n-m<
				return 
			} else {
				return 0
			}
		} else {
			continue
		}
	}
	return start
}

// @lc code=end
