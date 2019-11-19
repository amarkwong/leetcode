/*
 * @lc app=leetcode id=553 lang=golang
 *
 * [553] Optimal Division
 */

// @lc code=start
import "strconv"

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	} else if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	} else {
		res := strconv.Itoa(nums[0]) + "/(" + strconv.Itoa(nums[1])
		for i := 2; i < len(nums); i++ {
			res += "/" + strconv.Itoa(nums[i])
		}
		res += ")"
		return res
	}
}

// @lc code=end
