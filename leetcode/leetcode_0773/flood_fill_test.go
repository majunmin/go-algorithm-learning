/**
* Date: 2021/6/6 10:43 上午
* Author: majunmin
**/
package leetcode_0773

import "testing"

func Test_FlooadFill(t *testing.T) {

	temp1 := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	t.Log(floodFill(temp1, 1, 1, 2))
}
