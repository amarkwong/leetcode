/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 *
 * https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (40.47%)
 * Likes:    2500
 * Dislikes: 149
 * Total Accepted:    343.1K
 * Total Submissions: 846.5K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n1'
 *
 * Given a binary tree, find the lowest common ancestor (LCA) of two given
 * nodes in the tree.
 * 
 * According to the definition of LCA on Wikipedia: “The lowest common ancestor
 * is defined between two nodes p and q as the lowest node in T that has both p
 * and q as descendants (where we allow a node to be a descendant of itself).”
 * 
 * Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * Output: 3
 * Explanation: The LCA of nodes 5 and 1 is 3.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * Output: 5
 * Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant
 * of itself according to the LCA definition.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * All of the nodes' values will be unique.
 * p and q are different and both values will exist in the binary tree.
 * 
 * 
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func dfs(root, target *TreeNode) []*TreeNode {
	stack:=make([]*TreeNode)
	//top to bottom
	for ;n.Left!=nil||n.Right!=nil;{
		if n.Left==target||n.Right==target {
			return stack
		}else if n.Left!=nil{
			n=n.Left
			stack=append(stack,n)
		}else if n.Right!=nil{
			n=n.Right
			stack=append(stack,n)
		}
	}
	//bottom to root
	//pop the bottom nod
	stack=stack[:len(stack)-1]
	node:=stack[len(stack)-1:]

}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	stackp:=dfs(root,p)
	stackq:=dfs(root,q)
}
// @lc code=end

