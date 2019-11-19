import "math"

/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 */

// @lc code=start
//hashmap
// func findErrorNums(nums []int) []int {
// m := make(map[int]int)
// res := make([]int, 2)
// for idx, num := range nums {
// 	if _, ok := m[num]; ok {
// 		res[0] = num
// 	}
// 	m[num] = idx
// }

// for index := 1; index <= len(nums); index++ {
// 	if _, ok := m[index]; !ok {
// 		res[1] = index
// 		return res
// 	}
// }
// return res
// }
//flipping
func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	for _, num := range nums {
		if nums[int(math.Abs(float64(num)))-1] < 0 {
			res[0] = int(math.Abs(float64(num)))
		} else {
			nums[int(math.Abs(float64(num)))-1] *= -1
		}
	}
	for index := 1; index <= len(nums); index++ {
		if nums[index-1] > 0 {
			res[1] = index
			return res
		}
	}
	return res
}

// @lc code=end
