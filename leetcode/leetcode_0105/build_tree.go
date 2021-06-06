/**
* Date: 2021/6/2 7:46 上午
* Author: majunmin
**/
package leetcode_0105

import (
	. "go-algorithm-learning/leetcode/common"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeHelper(preorder, inorder)
}

// 递归构建` 二叉树
func buildTreeHelper(preorder []int, inorder []int) *TreeNode {
	// terminate cond
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := NewTreeNode(rootVal)
	var inRootIdx int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			inRootIdx = i
			break
		}
	}
	root.Left = buildTreeHelper(preorder[1:1+inRootIdx], inorder[:inRootIdx])
	root.Right = buildTreeHelper(preorder[1+inRootIdx:], inorder[inRootIdx+1:])

	return root
}
