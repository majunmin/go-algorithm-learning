/**
* Date: 2021/5/7 12:04 上午
* Author: majunmin
**/
package leetcode_0021

import "go-algorithm-learning/leetcode/common"

/**
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/
 * 1. 合并两个 有序链表  2. 合并两个有序数组
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	return dfsFunc(l1, l2)
}

func dfsFunc(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = dfsFunc(l1.Next, l2)
		return l1
	} else {
		l2.Next = dfsFunc(l1, l2.Next)
		return l2
	}
}

func iterFunc(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var dummyNode common.ListNode
	prev := &dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	// 合并后 l1 和 l2 最多只有一个还未被合并完，我们直接将链表末尾指向未合并完的链表即可
	if l1 != nil {
		prev.Next = l1
	}

	if l2 != nil {
		prev.Next = l2
	}
	return dummyNode.Next
}
