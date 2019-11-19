/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	row := []int{}
	queue := make([]*TreeNode, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for i := 0; len(queue) > 0; i++ {
		for _, node := range queue {
			row = append(row, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		res, row = append(res, row), nil
	}
	return res
}

// @lc code=end
