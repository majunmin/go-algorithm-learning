/**
  @Author: majm@ushareit.com
  @date: 2021/3/11
  @note:
**/
package tree

import "go-algorithm-learning/leetcode/common"

func postorderTraversal(root *common.TreeNode) []int {
	var stack []*common.TreeNode
	var result []int
	node := root
	var prev *common.TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if node.Right != nil && node.Right != prev {
			node = node.Right
			continue
		}
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		prev = node
		node = nil
	}
	return result
}
