/**
* Date: 2021/5/16 10:57 上午
* Author: majunmin
**/
package leetcode_0144

import (
	"go-algorithm-learning/leetcode/common"
	"testing"
)

func Test_PreOrder(t *testing.T) {
	root := common.TreeNode{
		Val:   1,
		Left:  &common.TreeNode{Val: 2},
		Right: &common.TreeNode{Val: 3},
	}
	preorderTraversal(&root)
}
