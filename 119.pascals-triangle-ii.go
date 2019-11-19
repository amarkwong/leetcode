/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
// @lc code=start
func getRow(rowIndex int) []int {
	S := make([]int, rowIndex+1)
	S[0] = 1
	for i := 1; i <= (rowIndex+1)/2; i++ {
		S[i] = S[i-1] * (rowIndex + 1 - i) / i
	}
	for i := rowIndex; i > (rowIndex+1)/2; i-- {
		S[i] = S[rowIndex-i]
	}
	return S
}

// @lc code=end
