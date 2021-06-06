/**
* Date: 2021/5/30 12:52 下午
* Author: majunmin
**/
package repeat_exercise

import "testing"

func Test_LongestSubString(t *testing.T) {
	// 最长字符 出现在 前面
	// 最长字符 出现在 中间
	// 最长字符 出现在 最后
	t.Log(lengthOfLongestSubstring("qbc"))
	t.Log(lengthOfLongestSubstring("abcfghcjkfabc"))
	t.Log(lengthOfLongestSubstring("dvdf"))
}
