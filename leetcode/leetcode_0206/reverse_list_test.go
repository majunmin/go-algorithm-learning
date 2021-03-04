/**
  @Author: majm@ushareit.com
  @date: 2021/2/25
  @note:
**/
package leetcode_0206

import (
	"fmt"
	. "go-algorithm-learning/leetcode/common"
	"testing"
)

func TestReverseList(t *testing.T) {
	node1 := ListNode{
		Next: nil,
		Val:  1,
	}
	node2 := ListNode{
		Next: nil,
		Val:  2,
	}
	node3 := ListNode{
		Next: nil,
		Val:  3,
	}
	node4 := ListNode{
		Next: nil,
		Val:  4,
	}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	rList := reverseList(&node1)
	fmt.Println(rList)
}
