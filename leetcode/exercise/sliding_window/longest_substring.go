/**
  @Author: majm@ushareit.com
  @date: 2021/3/16
  @note:
**/
package sliding_window

// 1. 窗口里面 肯定都是不重复的
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	right,result := 0,0
	set := make(map[byte]int)

	for i := range s {
		if i != 0 {
			delete(set, s[i-1])
		}
		for right < length && set[s[right]] == 0 {
			set[s[right]]++
			right++
			if right - i > result {
				result = right - i
			}
		}
	}
	return result
}
