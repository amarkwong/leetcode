/*
 * @lc app=leetcode id=532 lang=golang
 *
 * [532] K-diff Pairs in an Array
 */

// @lc code=start

func findPairs(nums []int, k int) int {
	m := make(map[int]int)
	n := make(map[int]int)
	cnt := 0
	if k < 0 {
		return 0
	}
	if k == 0 {
		for _, num := range nums {
			if _, ok := m[num]; ok {
				if m[num] == 1 {
					m[num]++
					cnt++
				}
			} else {
				m[num] = 1
			}
		}
	} else {
		for idx, num := range nums {
			m[num+k] = idx
			n[num] = idx
		}
		for idx, _ := range n {
			if _, ok := m[idx]; ok {
				cnt++
			}
		}
	}
	return cnt
}

// @lc code=end
