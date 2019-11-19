/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (36.72%)
 * Likes:    3147
 * Dislikes: 207
 * Total Accepted:    480.6K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 *
 * Example:
 *
 *
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getMin(heap []*ListNode) *ListNode{
	min=heap[0].Val
	for _,node:=heap[1:]{
		if node.Val<
	}
}
func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := make([]*ListNode, len(lists))
	head := make([]*ListNode, len(lists))
}

// @lc code=end
