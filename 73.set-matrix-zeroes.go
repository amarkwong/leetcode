/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	height, width := len(matrix), len(matrix[0])
	if height == 0 || width == 0 {
		return
	}
	hFlag, vFlag := make([]bool, width), make([]bool, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if matrix[i][j] == 0 {
				hFlag[j], vFlag[i] = true, true
			}
		}
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if hFlag[j] || vFlag[i] {
				matrix[i][j] = 0
			}
		}
	}
}

// @lc code=end
