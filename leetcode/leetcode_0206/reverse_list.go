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
	return dfsReverse1(head)

}

// 时间复杂度 O(N)
// 空间复杂度 O(N)
func dfsReverse1(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	next := dfsReverse1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}

// 迭代方式
// 时间复杂度 O(N)
// 空间复杂度 O(1)
func iterSolution1(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
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
		next := curNode.Next
		curNode.Next = prev
		prev = curNode
		curNode = next
	}
	return prev
}
