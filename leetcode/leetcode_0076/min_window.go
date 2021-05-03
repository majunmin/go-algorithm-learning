/**
  @Author: majm@ushareit.com
  @date: 2021/3/15
  @note:
**/
package leetcode_0076

import "math"

func minWindow(s string, t string) string {
	length := len(s)
	if length < len(t) {
		return ""
	}
	window, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	validCnt, validLen := 0, len(need)

	start, Len := -1, math.MaxInt32
	for left, right := 0, 0; right < length; {
		// 移动窗口
		u := s[right]
		if need[u] > 0 {
			window[u]++
			if need[u] == window[u] {
				validCnt++
			}
		}
		right++

		for validCnt == validLen {
			d := s[left]
			if right-left < Len {
				start = left
				Len = right - left
			}
			left++

			if need[d] > 0 && window[d] == need[d] {
				validCnt--
			}
			window[d]--
		}
	}

	return s[start : start+Len]
}
