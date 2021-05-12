/**
* Date: 2021/5/8 8:06 下午
* Author: majunmin
**/
package leetcode_0226

import "go-algorithm-learning/leetcode/common"

//recursive
//https://leetcode-cn.com/problems/invert-binary-tree/
func invertTree(root *common.TreeNode) *common.TreeNode {
	// terminate condition
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
