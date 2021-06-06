/**
* Date: 2021/5/27 8:27 上午
* Author: majunmin
**/
package leetcode_0019

import (
	. "go-algorithm-learning/leetcode/common"
)

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// 遍历两次 是肯定可以解决的 单链表 删除 a 节点 需要找到 其前驱节点 prev.Next = prev.Next.Next, 如果是 头结点?
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	dummyNode := &ListNode{}
	dummyNode.Next = head

	p, q := dummyNode, dummyNode
	for i := n; i > 0; i-- {
		q = q.Next
	}
	for q.Next != nil {
		q = q.Next
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummyNode.Next
}
