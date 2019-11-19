/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		m[target-num] = idx
	}
	for idx, num := range nums {
		if val, ok := m[num]; ok && val != idx {
			return []int{idx, val}
		}
	}
	return nil
}

// @lc code=end
