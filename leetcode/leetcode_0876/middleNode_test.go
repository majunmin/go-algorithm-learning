/**
* Date: 2021/5/28 9:04 下午
* Author: majunmin
**/
package leetcode_0876

import (
	. "go-algorithm-learning/leetcode/common"
	"testing"
)

func Test_MiddleNode(t *testing.T) {

	ll := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	middleNode(ll)
}
