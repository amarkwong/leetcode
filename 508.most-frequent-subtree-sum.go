/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func computeSubSum(root *TreeNode) {
	if root.Left != nil && root.Right != nil {
		computeSubSum(root.Left)
		computeSubSum(root.Right)
		root.Val = root.Left.Val + root.Right.Val + root.Val
	} else if root.Left != nil {
		computeSubSum(root.Left)
		root.Val = root.Left.Val + root.Val
	} else if root.Right != nil {
		computeSubSum(root.Right)
		root.Val = root.Right.Val + root.Val
	}
}
func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	m := make(map[int]int)
	res := []int{}
	computeSubSum(root)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		for _, node := range queue {
			m[node.Val]++
			if node.Left != nil {
				queue = append(queue, queue.Left)
			}
			if node.Right != nil {
				queue = append(queue, queue.Right)
			}
			queue = queue[1:]
		}
	}
}

// @lc code=end
