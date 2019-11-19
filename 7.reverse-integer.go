/* * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	num, res := x, 0
	for i := 0; num != 0; i++ {
		if res > 214748364 || res < -214748364 {
			return 0
		} else if res == 214748364 || res == -214748364 {
			if num%10 >= 7 {
				return 0
			} else {
				res = res*10 + num%10
				num = num / 10
			}
		} else {
			res = res*10 + num%10
			num = num / 10
		}
	}
	return res
}

// @lc code=end
