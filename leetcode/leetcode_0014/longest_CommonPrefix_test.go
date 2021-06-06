/**
* Date: 2021/6/6 5:20 下午
* Author: majunmin
**/
package leetcode_0014

import "testing"

func Test_LongestCommonPrefix(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	t.Log(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
