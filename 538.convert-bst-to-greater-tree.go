/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
 *
 * https://leetcode.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (52.70%)
 * Likes:    1525
 * Dislikes: 98
 * Total Accepted:    97.4K
 * Total Submissions: 184.7K
 * Testcase Example:  '[5,2,13]'
 *
 * Given a Binary Search Tree (BST), convert it to a Greater Tree such that
 * every key of the original BST is changed to the original key plus sum of all
 * keys greater than the original key in BST.
 *
 *
 * Example:
 *
 * Input: The root of a Binary Search Tree like this:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * Output: The root of a Greater Tree like this:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
 *
 *
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
import "fmt"

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Right)
	if root.Right != nil {
		root.Val += root.Right.Val
	}
	fmt.Println("Current node value", root.Val)
	inorder(root.Left)
}
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	inorder(root)
	return nil
}

// @lc code=end
