/**
  @Author: majm@ushareit.com
  @date: 2021/1/23
  @note: https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
**/

package leetcode_0104

import "go-algorithm-learning/leetcode/common"

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	lMaxDpeth := maxDepth(root.Left)
	rMaxDpeth := maxDepth(root.Right)
	return maxInt(lMaxDpeth, rMaxDpeth)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
