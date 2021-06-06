/**
* Date: 2021/5/16 11:38 上午
* Author: majunmin
**/
package leetcode_0145

import (
	"go-algorithm-learning/leetcode/common"
	"testing"
)

func Test_PostOrder(t *testing.T) {
	root := common.TreeNode{
		Val:   1,
		Left:  &common.TreeNode{Val: 2},
		Right: &common.TreeNode{Val: 3},
	}
	postorderTraversal(&root)
}
