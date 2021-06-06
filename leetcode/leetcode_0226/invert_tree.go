/**
* Date: 2021/5/8 8:06 下午
* Author: majunmin
**/
package leetcode_0226

import "go-algorithm-learning/leetcode/common"

//recursive
//https://leetcode-cn.com/problems/invert-binary-tree/
func invertTree(root *common.TreeNode) *common.TreeNode {
	return recurisiveSolution(root)
}

func recurisiveSolution(root *common.TreeNode) *common.TreeNode {
	// terminate
	if root == nil {
		return nil
	}
	//  process invert
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}
