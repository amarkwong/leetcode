/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	//denote the len(nums) as l
	//the subset will have the length of n=2^l
	//so just consider a n bit number, where
	//0...000 as the empty set while 1...111 as the
	//full set
	//now what we need to do is to map the digit into
	//the element/item in the nums array
	//the intuition is, for every element in the array
	//choose the start point from index 0 to l-1
	//and if that element is gonna to show in the subset
	//it index will be 1, otherwise will be 0
	l := uint32(len(nums))
	n := 2 << (l - 1)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		k := countBits(i)
		res[i] = make([]int, k)
		for j, m, n := i, 0, k; j != 0; j, m = j>>1, m+1 {
			if (n > 0) && (j%2 == 1) {
				res[i][k-n] = nums[m]
				n--
			}
		}
	}
	return res
}

func countBits(i int) int {
	count := 0
	for ; i != 0; count++ {
		i &= (i - 1)
	}
	return count
}

// @lc code=end
