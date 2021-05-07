/**
* Date: 2021/5/4 4:21 下午
* Author: majunmin
**/
package leetcode_0234

import "go-algorithm-learning/leetcode/common"

// 回文链表   --
// - 快慢指针法 画图一下会更好理解
// https://leetcode-cn.com/problems/palindrome-linked-list/
func isPalindrome(head *common.ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow, fast, prev := head, head, head
	var pprev *common.ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next

		prev.Next = pprev
		pprev = prev
	}

	if fast != nil {
		slow = slow.Next
	}
	for prev != nil && slow != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}

	return true
}
