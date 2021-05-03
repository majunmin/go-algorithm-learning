/**
  @Author: majm@ushareit.com
  @date: 2021/3/10
  @note:
**/
package leetcode_0094

import "go-algorithm-learning/leetcode/common"

func inorderTraversal(root *common.TreeNode) []int {
	return inorderIter(root)
}

func inorderIter(root *common.TreeNode) []int {

	var stack []*common.TreeNode
	var result []int
	node := root

	for node != nil || len(stack) > 0{
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// 出栈操作
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}

// 递归算法   简单易懂
func inorder(root *common.TreeNode) []int {
	var result []int
	inorderTraversalHelper(root, &result)
	return result
}

func inorderTraversalHelper(root *common.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorderTraversalHelper(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraversalHelper(root.Right, result)
}
