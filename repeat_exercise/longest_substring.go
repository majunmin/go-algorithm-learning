/**
* Date: 2021/5/30 12:07 下午
* Author: majunmin
**/
package repeat_exercise

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {

	// 滑动窗口法  基础
	set := make(map[byte]int)
	l, r := 0, 0
	maxLen := 0
	for r < len(s) {
		// 如果没有重复字符 移动右指针
		if _, exist := set[s[r]]; !exist {
			set[s[r]] = r
			r++
			maxLen = maxInt(maxLen, r-l)
			continue
		}

		// 如果有重复字符 移动左指针

		for l <= set[s[r]] {
			delete(set, s[l])
			l++
		}

	}
	return maxLen
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
