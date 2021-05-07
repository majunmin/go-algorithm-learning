/**
* Date: 2021/5/7 12:11 上午
* Author: majunmin
**/
package leetcode_0021

import (
	"go-algorithm-learning/leetcode/common"
	"testing"
)

func TestMergeTowList(t *testing.T) {
	l1 := common.ListNode{Val: 1}
	l2 := common.ListNode{Val: 2}
	l3 := common.ListNode{Val: 4}

	l4 := common.ListNode{Val: 1}
	l5 := common.ListNode{Val: 3}
	l6 := common.ListNode{Val: 4}

	l1.Next, l2.Next = &l2, &l3
	l4.Next, l5.Next = &l5, &l6

	result := mergeTwoLists(&l1, &l4)
	common.DisPlayListNode(result)

}
