/**
* Date: 2021/5/28 9:30 下午
* Author: majunmin
**/
package leetcode_0002

import (
	"fmt"
	. "go-algorithm-learning/leetcode/common"
	"testing"
)

func Test_towSum(t *testing.T) {
	l1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l5 := &ListNode{
		Val:  6,
		Next: nil,
	}
	l6 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l1.Next = l2
	l2.Next = l3
	//l3.Next = l4
	l4.Next = l5
	l5.Next = l6

	numbers := addTwoNumbers(l1, l4)
	fmt.Println(numbers)
}
