/**
* Date: 2021/5/29 8:00 上午
* Author: majunmin
**/
package leetcode_0011

import "testing"

func Test_MaxArea(t *testing.T) {
	var a []int
	a = []int{2, 3, 4, 5, 18, 17, 6}
	t.Log(maxArea(a))

	a = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(a))

	a = []int{1, 2, 1}
	t.Log(maxArea(a))

	a = []int{1, 1}
	t.Log(maxArea(a))

	a = []int{4, 3, 2, 1, 4}
	t.Log(maxArea(a))
}
