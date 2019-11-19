/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
func thirdMax(nums []int) int {
	m := make(map[int]int)
	max1, max2, max3 := -2147483648, -2147483648, -2147483648
	for k, v := range nums {
		m[v] = k
		if v == max1 || v == max2 || v == max3 {
			continue
		}
		if v > max1 {
			max1, max2, max3 = v, max1, max2
		} else if v > max2 {
			max2, max3 = v, max2
		} else if v > max3 {
			max3 = v
		}
	}
	if len(m) >= 3 {
		return max3
	} else {
		return max1
	}
}

// @lc code=end
