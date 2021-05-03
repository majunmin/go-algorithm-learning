/**
  @Author: majm@ushareit.com
  @date: 2021/3/11
  @note:
**/
package tree

import "go-algorithm-learning/leetcode/common"

func preorderTraversal(root *common.TreeNode) []int {
	return preorderIter(root)
}

func preorderIter(root *common.TreeNode) []int {
	var stack []*common.TreeNode
	var result []int
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			result = append(result, node.Val)
			node = node.Left
		}
		// 出栈操作
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return result
}