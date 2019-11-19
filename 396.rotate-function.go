/*
 * @lc app=leetcode id=396 lang=golang
 *
 * [396] Rotate Function
 */

// @lc code=start
func maxRotateFunction(A []int) int {
	sum := 0
	f := 0
	n := len(A)
	for idx, num := range A {
		sum += num
		f += idx * num
	}
	max := f
	for i := 1; i < n; i++ {
		f += sum - n*A[n-i]
		if f > max {
			max = f
		}
	}
	return max
}

// @lc code=end
