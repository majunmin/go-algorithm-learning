/**
* Date: 2021/5/4 3:49 下午
* Author: majunmin
**/
package leetcode_0125

import "unicode"

// 验证回文子串
// https://leetcode-cn.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
	return isPalindrome2(s)
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func isPalindrome2(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAOrNum(rune(s[left])) {
			left++
		}

		for left < right && !isAOrNum(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

// common 1. 先提取出合法的字符串 2. 在进行验证
// 时间复杂度O(n)
// 空间复杂度 O(n)
func common(s string) bool {
	var sgood string
	for _, c := range s {
		if isAOrNum(c) {
			sgood += string(unicode.ToLower(c))
		}
	}
	length, mid := len(sgood), len(sgood)/2

	for i := 0; i < mid; i++ {
		if sgood[i] != sgood[length-1-i] {
			return false
		}
	}
	return true
}

func isAOrNum(c int32) bool {
	return unicode.IsLetter(c) || unicode.IsNumber(c)
}
