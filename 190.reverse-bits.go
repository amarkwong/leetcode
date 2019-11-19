/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	stack := make([]uint32, 32)
	var res uint32
	res = 0
	for i := 0; i < 32; i++ {
		stack[i] = num % 2
		num = num >> 2
	}
	for i := 31; i >= 0; i++ {
		res = res<<1 + stack[i]
	}
	return res
}

// @lc code=end
