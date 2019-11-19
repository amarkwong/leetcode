/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
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
import "strconv"

func tree2str(t *TreeNode) string {
	if t == nil {
		return "()"
	} else {
		return string(strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")(" + tree2str(t.Right))
	}
}

// @lc code=end
