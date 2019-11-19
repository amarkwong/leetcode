/*
 * @lc app=leetcode id=517 lang=golang
 *
 * [517] Super Washing Machines
 */

// @lc code=start
func findMinMoves(machines []int) int {
	for i,dress:=0,0;i<len(machines);i++
		dress+=machines[i]
	}
	if (dress%len(machines)!=0){
		return -1
	}else{
		
	}
}
// @lc code=end

