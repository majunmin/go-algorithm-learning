/**
* Date: 2021/5/29 9:26 上午
* Author: majunmin
**/
package leetcode_0020

// https://leetcode-cn.com/problems/valid-parentheses/
//
func isValid(s string) bool {
	validParentheses := map[byte]byte{'}': '{', ')': '(', ']': '['}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		_, exist := validParentheses[s[i]]
		if !exist {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 || validParentheses[s[i]] != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
