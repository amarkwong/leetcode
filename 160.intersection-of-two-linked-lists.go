/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	na, nb := headA, headB
	for na != nil || nb != nil {
		if na != nil {
			m[na]++
		}
		if nb != nil {
			m[nb]++
		}
		if m[na] > 1 {
			return na
		} else if m[nb] > 1 {
			return nb
		}
		if na != nil {
			na = na.Next
		}
		if nb != nil {
			nb = nb.Next
		}
	}
	return nil
}

// @lc code=end
