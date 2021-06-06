/**
* Date: 2021/5/28 8:57 下午
* Author: majunmin
**/
package leetcode_0876

import (
	. "go-algorithm-learning/leetcode/common"
)

// https://leetcode-cn.com/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	return slow
}
