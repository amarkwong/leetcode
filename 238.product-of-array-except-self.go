import "fmt"

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=starto
func productExceptSelf(nums []int) []int {
	prod := 1
	zeroCount := 0
	index := 0
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
		if nums[i] == 0 && zeroCount == 0 {
			index = i
			zeroCount++
		} else {
			if nums[i] == 0 && zeroCount == 1 {
				zeroCount++
				for j := 0; j < len(nums); j++ {
					result[j] = 0
				}
				return result
			} else {
				prod *= nums[i]
			}
		}
	}

	if zeroCount == 1 {
		for j := 0; j < len(nums); j++ {
			if j == index {
				result[j] = prod
			} else {
				result[j] = 0
			}
		}
		return result
	} else {
		for j := 0; j < len(nums); j++ {
			result[j] = prod / nums[j]
		}
		return result
	}
}

// @lc code=end
