import "fmt"

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	s := nums[:]
	s1 := nums[:l-k]
	s2 := nums[l-k:]
	fmt.Println(s1, s2)
	s = append(s2, s1...)
	for i := 0; i < l; i++ {
		nums[i] = s[i]
	}
}

// @lc code=end
