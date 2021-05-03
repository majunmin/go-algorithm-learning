/**
  @Author: majm@ushareit.com
  @date: 2021/3/11
  @note:
**/
package tree

import (
	"fmt"
	"go-algorithm-learning/leetcode/common"
	"testing"
)

func TestTravelsal(t *testing.T) {
	node3 := &common.TreeNode{
		Val: 3,
	}
	node2 := &common.TreeNode{
		Val:  2,
		Left: node3,
	}
	root := common.TreeNode{
		Val:   1,
		Right: node2,
	}

	fmt.Println(postorderTraversal(&root))
}
