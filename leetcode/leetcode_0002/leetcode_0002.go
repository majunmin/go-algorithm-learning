/**
* Date: 2021/5/28 9:23 下午
* Author: majunmin
**/
package leetcode_0002

import (
	. "go-algorithm-learning/leetcode/common"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	ll := dummyNode
	var i int
	for l1 != nil || l2 != nil {
		node := &ListNode{}
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		node.Val = (v1 + v2 + i) % 10
		i = (v1 + v2 + i) / 10
		ll.Next = node
		ll = ll.Next
	}
	if i != 0 {
		ll.Next = &ListNode{Val: i}
	}
	return dummyNode.Next
}
