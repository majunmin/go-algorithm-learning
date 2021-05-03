/**
  @Author: majm@ushareit.com
  @date: 2021/3/15
  @note:
**/
package leetcode_0567

func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len > s2Len {
		return false
	}

	var cnt1, cnt2 [26]int
	for i := 0; i < s1Len; i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}

	// 数组的判等 学了一招呀
	// 判断字符串是否含有的 子串字符是否相等  变为数组判等 高啊
	if cnt1 == cnt2 {
		return true
	}

	for i := s1Len; i < len(s2); i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-s1Len]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}
