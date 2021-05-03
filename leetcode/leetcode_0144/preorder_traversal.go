/**
  @Author: majm@ushareit.com
  @date: 2021/3/10
  @note:
**/
package leetcode_0144

import (
	"go-algorithm-learning/leetcode/common"
)

func preorderTraversal(root *common.TreeNode) []int {
	return preorderIter(root)
}

// 解法2 迭代算法
func preorderIter(root *common.TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*common.TreeNode
	var result []int

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			result = append(result, node.Val)
			node = node.Left
		}

		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return result
}

// 递归算法 解法比较简单
func preorder(root *common.TreeNode) []int {
	var result []int
	preorderTraversalHelper(root, &result)
	return result
}

func preorderTraversalHelper(root *common.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	preorderTraversalHelper(root.Left, result)
	preorderTraversalHelper(root.Right, result)
}
