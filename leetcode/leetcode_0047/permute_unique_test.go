/**
* Date: 2021/6/6 8:46 上午
* Author: majunmin
**/
package leetcode_0047

import "testing"

func Test_PermuteUnique(t *testing.T) {
	t.Log(permuteUnique([]int{1, 1, 2}))
	t.Log(permuteUnique([]int{1, 2, 3}))
	t.Log(permuteUnique([]int{}))
}
