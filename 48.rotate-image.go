/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	size := len(matrix)
	for i := 0; i < size/2; i++ {
		matrix[i], matrix[size-1-i] = matrix[size-1-i], matrix[i]
	}
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// @lc code=end
