/**
  @Author: majm@ushareit.com
  @date: 2021/3/10
  @note:
**/
package leetcode_0110

import "go-algorithm-learning/leetcode/common"

type Result struct {
	isBalanced bool
	height     int
}

func isBalanced(root *common.TreeNode) bool {
	res := isBalancedHelp(root)
	return res.isBalanced
}

func isBalancedHelp(root *common.TreeNode) Result {
	if root == nil {
		return Result{
			isBalanced: true,
			height:     0,
		}
	}
	lRes := isBalancedHelp(root.Left)
	rRes := isBalancedHelp(root.Right)

	return Result{
		isBalanced: lRes.isBalanced && rRes.isBalanced && abs(lRes.height-rRes.height) <= 1,
		height:     maxInt(lRes.height, rRes.height) + 1,
	}
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

// 1. 递归解法
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
