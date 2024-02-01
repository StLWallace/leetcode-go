/* https://leetcode.com/problems/maximum-depth-of-binary-tree/?envType=study-plan-v2&envId=leetcode-75

Accepted:
Runtime - 4ms
Memory - 4.48MB
*/

package maxdepthbintree

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	max_depth := 0
	l_depth := 0
	r_depth := 0
	if root != nil {
		max_depth = 1
		if root.Left != nil {
			l_depth = maxDepth(root.Left)
		}
		if root.Right != nil {
			r_depth = maxDepth(root.Right)
		}
	}
	if l_depth > r_depth {
		max_depth += l_depth
	} else {
		max_depth += r_depth
	}
	return max_depth
}