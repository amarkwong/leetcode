/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start

func searchMatrix(matrix [][]int, target int) bool {
	//go through the diagonal line
	return helper(matrix, target, 0, 0, len(matrix), len(matrix[0]))
}
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return b
	}
}
func helper(matrix [][]int, target int, vOffset int, hOffset int, height int, width int) bool {
	dim := min(height, width)
	if dim <= 0 {
		return false
	}
	for i := 0; i < dim; i++ {
		if target == matrix[vOffset+i][hOffset+i] {
			return true
		} else if target < matrix[vOffset+i][hOffset+i] {
			return helper(matrix, target, vOffset, hOffset+i, i, width-i) || helper(matrix, target, vOffset+i, hOffset, height-i, i)
		} else if target > matrix[vOffset+i][hOffset+i] {
			continue
		}
	}
	if height == width {
		return false
	} else if height > width {
		return helper(matrix, target, vOffset+dim, hOffset, height-dim, dim)
	} else {
		return helper(matrix, target, vOffset, hOffset+dim, dim, width-dim)
	}
}

// @lc code=end
