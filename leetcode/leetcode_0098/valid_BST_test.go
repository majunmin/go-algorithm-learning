/**
* Date: 2021/5/16 12:59 上午
* Author: majunmin
**/
package leetcode_0098

import (
	"fmt"
	"go-algorithm-learning/leetcode/common"
	"testing"
)

func Test_IsValidBST(t *testing.T) {
	root := &common.TreeNode{
		Val: 5,
		Left: &common.TreeNode{
			Val: 4,
		},
		Right: &common.TreeNode{
			Val: 6,
			Left: &common.TreeNode{
				Val: 3,
			},
			Right: &common.TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(isValidBST(root))
}
