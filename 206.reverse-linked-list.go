/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	//iteratively
	//create a stack, push the current nod
	//pop when hit the tail
	if head == nil {
		return nil
	}
	s := make([]*ListNode, 0)
	for n := head; n != nil; n = n.Next {
		s = append(s, n)
	}
	l := len(s) - 1
	h := s[l]
	for res := s[l]; l >= 0; l-- {
		if l == 0 {
			res.Next = nil
		} else {
			res.Next = s[l-1]
			res = res.Next
		}
	}
	return h
}

// @lc code=end
