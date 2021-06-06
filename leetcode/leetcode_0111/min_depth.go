/**
* Date: 2021/5/16 5:11 下午
* Author: majunmin
**/
package leetcode_0111

import (
	"go-algorithm-learning/leetcode/common"
	"math"
)

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func minnDepth(root *common.TreeNode) int {

	return solution_iter(root)

}

// 广度优先搜索
// 时间复杂度 o(N)
// 空间复杂度 O(N)
func solution_iter(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	minDep := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return minDep + 1
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		minDep += 1
	}
	return minDep
}

// 递归解法
// 时间复杂度   O(N)
// 空间复杂度   树的高度  O(H)  ==> O(logN)
func solution_recur(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	// 是叶子节点  ==> 1
	if root.Left == nil && root.Right == nil {
		return 1
	}

	minDep := math.MaxInt32
	if root.Left != nil {
		minDep = minInt(solution_recur(root.Left), minDep)
	}
	if root.Right != nil {
		minDep = minInt(solution_recur(root.Right), minDep)
	}
	return minDep + 1
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
