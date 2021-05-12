/**
  @Author: majm@ushareit.com
  @date: 2021/3/10
  @note:
**/
package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
