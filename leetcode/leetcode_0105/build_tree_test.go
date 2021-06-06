/**
* Date: 2021/6/2 7:57 上午
* Author: majunmin
**/
package leetcode_0105

import (
	"fmt"
	"testing"
)

func Test_BuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}
