/**
* Date: 2021/5/8 7:43 下午
* Author: majunmin
**/
package offer_055

import (
	"go-algorithm-learning/leetcode/common"
	"testing"
)

func Test_IsBalanced(t *testing.T) {
	root := common.NewTreeNode(1)
	l1 := common.NewTreeNode(2)
	r1 := common.NewTreeNode(2)
	l11 := common.NewTreeNode(3)
	r12 := common.NewTreeNode(3)
	l111 := common.NewTreeNode(3)
	r122 := common.NewTreeNode(3)
	//r121 := common.NewTreeNode(6)

	root.Left, root.Right = l1, r1
	l1.Left = l11
	r1.Right = r12
	l11.Left = l111
	r12.Right = r122
	//r12.Left = r121
	println(isBalanced(root))
}
