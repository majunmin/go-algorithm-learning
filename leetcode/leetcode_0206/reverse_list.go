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
	return dfsReverse(head)

}

func dfsReverse(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}

	newNode := dfsReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode

}

func iterSolution(head *ListNode) *ListNode {
	var newNode *ListNode
	curNode := head
	for curNode != nil {
		head = curNode.Next
		curNode.Next = newNode
		newNode = curNode

		curNode = head
	}
	return newNode
}
