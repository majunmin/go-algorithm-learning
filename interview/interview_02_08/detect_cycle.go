/**
* Date: 2021/5/27 8:05 上午
* Author: majunmin
**/
package interview_02_08

import . "go-algorithm-learning/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode-cn.com/problems/linked-list-cycle-lcci/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast, slow := head, head
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { // 之后在看题思考一下
			return fast
		}
	}
	return nil
}
