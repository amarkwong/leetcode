import "fmt"

/*
 * @lc app=leetcode id=725 lang=golang
 *
 * [725] Split Linked List in Parts
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	size := make([]int, k)
	res := make([]*ListNode, k)
	len := 0
	for node := root; node != nil; len++ {
		node = node.Next
	}
	if len%k == 0 {
		for index := 0; index < k; index++ {
			size[index] = len / k
		}
	} else {
		for index := 0; index < len%k; index++ {
			size[index] = len/k + 1
		}
		for index := len % k; index < k; index++ {
			size[index] = len / k
		}
	}
	fmt.Println("len :", len, " size: ", size)
	for i, h := 0, root; i < k; i++ {
		for j, l := 0, h; j < size[i]; j++ {
			res[i] = h
			if j == size[i]-1 {
				if l == nil {
					h = nil
				} else {
					h = l.Next
					l.Next = nil
				}
			} else {
				l = l.Next
			}
		}
	}
	return res
}

// @lc code=end

// @lc code=end
