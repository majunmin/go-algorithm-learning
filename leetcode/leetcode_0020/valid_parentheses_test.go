/**
* Date: 2021/5/29 9:39 上午
* Author: majunmin
**/
package leetcode_0020

import "testing"

func Test_IsValid(t *testing.T) {
	t.Log(isValid("()"))
	t.Log(isValid("([)]"))
	t.Log(isValid("[()]{}"))
}
