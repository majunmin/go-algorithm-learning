/**
  @Author: majm@ushareit.com
  @date: 2021/3/10
  @note:
**/
package leetcode_0145

import (
	"go-algorithm-learning/leetcode/common"
)

func postorderTraversal(root *common.TreeNode) []int {
	return postOrderIter(root)
}

func postOrderIter(root *common.TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var stack []*common.TreeNode
	var lastNode *common.TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if node.Right == nil || node.Right == lastNode {
			result = append(result, node.Val)
			lastNode = node
			node = nil
			continue
		}
		stack = append(stack, node)
		node = node.Right
	}
	return result
}

// 1. 递归解法
func postOrder(root *common.TreeNode) []int {
	var result []int
	postorderTraversalHelper(root, &result)
	return result
}

func postorderTraversalHelper(root *common.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postorderTraversalHelper(root.Left, result)
	postorderTraversalHelper(root.Right, result)
	*result = append(*result, root.Val)
}
