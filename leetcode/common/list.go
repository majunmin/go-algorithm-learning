/**
  @Author: majm@ushareit.com
  @date: 2021/2/25
  @note:
**/
package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func DisPlayListNode(node *ListNode) {
	var res []int
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	fmt.Println(res)
}
