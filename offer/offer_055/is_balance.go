/**
* Date: 2021/5/8 7:33 下午
* Author: majunmin
**/
package offer_055

import "go-algorithm-learning/leetcode/common"

// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
// 如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树
func isBalanced(root *common.TreeNode) bool {
	result := Balanced(root)
	return result.IsBalance
}

func Balanced(root *common.TreeNode) *Result {
	// terminate
	if root == nil {
		return NewResult(true, 0)
	}

	lRes := Balanced(root.Left)
	rRes := Balanced(root.Right)
	return NewResult(lRes.IsBalance && rRes.IsBalance && absInt(lRes.Height-rRes.Height) <= 1, 1+common.MaxInt(lRes.Height, rRes.Height))
}

type Result struct {
	IsBalance bool
	Height    int
}

func NewResult(isBalance bool, height int) *Result {
	return &Result{IsBalance: isBalance, Height: height}
}

func absInt(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
