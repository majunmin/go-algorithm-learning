/**
* Date: 2021/5/16 12:54 上午
* Author: majunmin
**/
package leetcode_0098

import (
	"go-algorithm-learning/leetcode/common"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode-cn.com/problems/validate-binary-search-tree/
func isValidBST(root *common.TreeNode) bool {
	return helper(root, math.MinInt32, math.MaxInt32)
}

// 时间复杂度  O(n)
// 空间复杂度  O(logn)
func helper(root *common.TreeNode, lower int, upper int) bool {
	// terminate
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
