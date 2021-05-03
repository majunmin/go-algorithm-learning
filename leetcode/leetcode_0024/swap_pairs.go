/**
  @Author: majm@ushareit.com
  @date: 2021/3/10
  @note:
**/
package leetcode_0024

import (
	"go-algorithm-learning/leetcode/common"
)

func swapPairs(head *common.ListNode) *common.ListNode {

	return swapPairsHelp(head)

}

func swapPairsHelp(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextNode := head.Next
	head.Next = swapPairs(nextNode.Next)
	nextNode.Next = head
	return nextNode
}

func iterSolution(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &common.ListNode{Next: head}
	tmp := dummyNode
	for tmp.Next != nil && tmp.Next.Next != nil {
		p := tmp.Next
		q := tmp.Next.Next
		tmp.Next = q
		p.Next = q.Next
		q.Next = p
		tmp = p
	}

	return dummyNode.Next
}
