/**
* Date: 2021/5/16 11:02 下午
* Author: majunmin
**/
package leetcode_0297

import (
	"go-algorithm-learning/leetcode/common"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *common.TreeNode) string {
	result := make([]string, 0)
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			// 出栈
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				result = append(result, "null")
				continue
			}
			result = append(result, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return "[" + strings.Join(result, ",") + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *common.TreeNode {
	if !strings.HasPrefix(data, "[") || !strings.HasSuffix(data, "]") {
		return nil
	}

	data = data[1 : len(data)-1]
	if len(data) == 0 {
		return nil
	}
	dataArr := strings.Split(data, ",")
	queue := make([]*common.TreeNode, 0)
	rootVal, _ := strconv.Atoi(dataArr[0])
	root := common.NewTreeNode(rootVal)
	queue = append(queue, root)

	idx := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if "null" != dataArr[idx] {
			val, _ := strconv.Atoi(dataArr[idx])
			node.Left = common.NewTreeNode(val)
			queue = append(queue, node.Left)
		}
		idx++
		if "null" != dataArr[idx] {
			val, _ := strconv.Atoi(dataArr[idx])
			node.Right = common.NewTreeNode(val)
			queue = append(queue, node.Right)
		}
		idx++
	}
	return root
}

func NewTreeNode(val int) *common.TreeNode {
	return &common.TreeNode{Val: val}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
