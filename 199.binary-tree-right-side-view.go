/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res, row := make([]int, 0), make([]int, 0)
	queue := make([]*TreeNode, 0)
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
		res, row = append(res, row[len(row)-1:]...), nil
	}
	return res
}

// @lc code=end
