/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carryDigit := false

	//consider the common case first
	newHead := new(ListNode)
	for p1, p2, n := l1, l2, newHead; p1 != nil || p2 != nil || carryDigit; {
		if p1 == nil && p2 != nil {
			if carryDigit {
				n.Val = p2.Val + 1
			} else {
				n.Val = p2.Val
			}
			carryDigit = (n.Val >= 10)
			if carryDigit {
				n.Val = n.Val - 10
			}
			if p2.Next != nil || carryDigit {
				n.Next = new(ListNode)
				n = n.Next
			}
			p2 = p2.Next
		} else if p1 != nil && p2 == nil {
			if carryDigit {
				n.Val = p1.Val + 1
			} else {
				n.Val = p1.Val
			}
			carryDigit = (n.Val >= 10)
			if carryDigit {
				n.Val = n.Val - 10
			}
			if p1.Next != nil || carryDigit {
				n.Next = new(ListNode)
				n = n.Next
			}
			p1 = p1.Next
		} else if p1 != nil && p2 != nil {
			if carryDigit {
				n.Val = p1.Val + p2.Val + 1
			} else {
				n.Val = p1.Val + p2.Val
			}
			carryDigit = (n.Val >= 10)
			if carryDigit {
				n.Val = n.Val - 10
			}
			if p1.Next != nil || p2.Next != nil || carryDigit {
				n.Next = new(ListNode)
				n = n.Next
			}
			p1 = p1.Next
			p2 = p2.Next
		} else {
			n.Val = 1
			carryDigit = false
		}
	}
	return newHead
}

// @lc code=end
