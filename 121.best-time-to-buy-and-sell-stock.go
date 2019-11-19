/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	//using dp, although you can only perform one
	//transaction, you can think that you are able
	//to buy one and sell one at the same day

	//handle exception
	if len(prices) <= 1 {
		return 0
	}
	maxSpread, curSpread := 0, 0
	for i := 1; i < len(prices); i++ {
		curSpread += prices[i] - prices[i-1]
		curSpread = max(0, curSpread)
		maxSpread = max(curSpread, maxSpread)
	}
	return maxSpread

}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
