/**
  @Author: majm@ushareit.com
  @date: 2021/2/24
  @note:
**/
package leetcode_0003

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	return solution2(s)
}

// 解法1的优化
// - 左指针每次仅仅移动一个, 会增加无谓的 迭代，可以将左指针移动到 第一个重复元素的 右边
// - map 记录指针位置
func solution2(s string) int {
	mp := make(map[byte]int)
	left, right, length := 0, -1, len(s)
	result := 0
	for left < length {
		for mp[s[right+1]] == 0 {
			mp[s[right+1]] = right + 1
			right++
		}

		if right-left+1 > result {
			result = right - left + 1
		}
		left++
	}
	return result
}

// 依次移动左指针  找到最大长度的不重复子串
// 时间复杂度 O(N)
// 空间复杂度 O(N)
func solution1(s string) int {
	set := make(map[byte]int8)
	sVec := []byte(s)

	rk, length, result := -1, len(s), 0
	for left, b := range sVec {
		for rk+1 < length && set[sVec[rk+1]] == 0 {
			set[sVec[rk+1]] = 1
			rk++
		}
		if rk-left+1 > result {
			result = rk - left + 1
		}
		delete(set, b)
	}
	return result
}
