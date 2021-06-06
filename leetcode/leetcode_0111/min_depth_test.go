/**
* Date: 2021/5/16 5:38 下午
* Author: majunmin
**/
package leetcode_0111

import (
	"fmt"
	"go-algorithm-learning/leetcode/common"
	"testing"
)

func Test_MinDepth(t *testing.T) {
	root := common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:   2,
			Left:  &common.TreeNode{Val: 4},
			Right: &common.TreeNode{Val: 5},
		},
		Right: &common.TreeNode{
			Val:   3,
			Left:  &common.TreeNode{Val: 6},
			Right: &common.TreeNode{Val: 7},
		},
	}
	fmt.Println(minnDepth(&root))

}
