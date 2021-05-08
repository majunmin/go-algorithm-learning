/**
* Date: 2021/5/8 9:00 上午
* Author: majunmin
**/
package tree

import "go-algorithm-learning/leetcode/common"

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 1
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
		maxDepth++
	}
	return maxDepth
}
