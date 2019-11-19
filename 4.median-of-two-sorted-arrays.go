/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//odd count
	if len(nums1)+len(nums2)%2==1{
		if len(nums1)==0 {
			return float64(nums2[len(nums2)/2])
		}
		if len(nums2)==0 {
			return float64(nums1[len(nums1)/2])
		}
		for 
	}else{
	//even count
	}
}
// @lc code=end

