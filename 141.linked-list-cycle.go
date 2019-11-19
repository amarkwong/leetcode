/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)
	if head == nil {
		return false
	}
	for p, i := head, 1; p.Next != nil; i++ {
		if _, ok := m[p.Next]; ok {
			return true
		} else {
			m[p] = i
			p = p.Next
		}
	}
	return false
}

// @lc code=end
