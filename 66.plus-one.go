/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	res := make([]int, len(digits))
	copy(res, digits)
	if len(res) == 0 {
		res = append([]int{1}, res...)
		return res
	}
	//plus one
	res[len(res)-1]++
	//handle the carry digit
	for i := len(res) - 1; res[i] == 10 && i >= 0; i-- {
		res[i] = 0
		if i > 0 {
			res[i-1]++
		} else {
			res = append([]int{1}, res...)
			return res
		}
	}
	return res
}

// @lc code=end
