/**
* Date: 2021/5/16 11:03 下午
* Author: majunmin
**/
package leetcode_0297

import (
	"fmt"
	"go-algorithm-learning/leetcode/common"
	"reflect"
	"testing"
)

func Test_Codec(t *testing.T) {
	//root := common.TreeNode{
	//	Val:   1,
	//	Left:  &common.TreeNode{Val: 2},
	//	Right: &common.TreeNode{
	//		Val:   3,
	//		Left:  &common.TreeNode{Val: 4},
	//		Right: &common.TreeNode{Val: 5},
	//	},
	//}

	//ser := Constructor()
	deser := Constructor()
	//data := ser.serialize(&root)
	//ans := deser.deserialize(data)
	ans1 := deser.deserialize("[]")
	fmt.Println(ans1)
}

func TestCodec_deserialize(t *testing.T) {

}

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Codec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTreeNode(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want *common.TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTreeNode(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
