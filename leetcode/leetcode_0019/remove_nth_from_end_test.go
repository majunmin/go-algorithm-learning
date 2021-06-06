/**
* Date: 2021/5/27 8:38 上午
* Author: majunmin
**/
package leetcode_0019

import (
	"fmt"
	"testing"

	. "go-algorithm-learning/leetcode/common"
)

func Test_removeNthFromEnd(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	removeNthFromEnd(l1, 5)
	fmt.Println(l1)
}
