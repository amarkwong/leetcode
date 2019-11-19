/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//regular basis
	h := l1
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2, h = l2, l1, l2
	}
	for curList := true; l1 != nil && l2 != nil; {
		if curList {
			if l1.Next == nil || l1.Next.Val > l2.Val {
				l1.Next, l1, curList = l2, l1.Next, false
			} else {
				l1 = l1.Next
			}
		} else {
			if l2.Next == nil || l2.Next.Val > l1.Val {
				l2.Next, l2, curList = l1, l2.Next, true
			} else {
				l2 = l2.Next
			}
		}
	}
	return h
}

// @lc code=end
