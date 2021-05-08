/**
  @Author: majm@ushareit.com
  @date: 2021/2/25
  @note:
**/
package leetcode_0206

import (
	. "go-algorithm-learning/leetcode/common"
)

// https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	return iterSolution(head)

}

func dfsReverse(head *ListNode) *ListNode {

	// terminate condition
	if head == nil || head.Next == nil {
		return head
	}

	// process
	next := dfsReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next

}

func iterSolution(head *ListNode) *ListNode {
	var prev *ListNode
	curNode := head
	for curNode != nil {
		head = curNode.Next
		curNode.Next = prev
		prev = curNode
		curNode = head
	}
	return prev
}
