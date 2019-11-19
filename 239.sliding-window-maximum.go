/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
func max(i int, n int) int {
	if i >= n {
		return i
	} else {
		return n
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	// init, get the max(max) and count of duplicate max(maxCount) in k, as well as the second max sm, second max count
	// smCount
	// check the first item i ,
	if i == max && maxCount > 1 {
		// compare i with new nubmer n
		if i > n {
			maxCount--
		}
		if i < n {
			max = n
			maxCount = 1
		}
	}
	if i == max && maxCount == 1 {
		// compare second max with new nubmer n
		if sm > n {
			max = sm
			maxCount = 1
		}
		if sm == n {
			max = sm
			maxCount = 2
		}
		if sm < n {
			max = n
			maxCount = 1
		}
	}
	if i != max {
		if max == n {
			maxCount++
		}
		if max < n {
			sm = max
			smCount = maxCount
			max = n
			maxCount = 1
		}
		if max > n {
			if n > sm {
				sm = n
				smCount = 1
			}
			if n == sm {
				smCount++
			}
		}
	}
}

// @lc code=end
