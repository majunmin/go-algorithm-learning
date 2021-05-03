/**
  @Author: majm@ushareit.com
  @date: 2021/3/15
  @note:
**/
package leetcode_0438

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
// 找到字符串中所有字母异位词
// 这道题和我想的 不太一样   abc 和 abc 也算异位词了
func findAnagrams(s string, p string) []int {
	return func2(s, p)
}

func func2(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	var result []int
	var validCnt int
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	for left, right := 0, 0; right < len(s); {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				validCnt++
			}
		}

		if right-left >= len(p) {
			if validCnt == len(need) {
				result = append(result, left)
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					validCnt--
				}
				window[d]--
			}
		}
	}

	return result
}

func func1(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	var result []int

	var windowCnt, targetCnt [26]int
	for i := 0; i < pLen; i++ {
		windowCnt[s[i]-'a']++
		targetCnt[p[i]-'a']++
	}

	if windowCnt == targetCnt {
		result = append(result, 0)
	}

	for i := pLen; i < sLen; i++ {
		windowCnt[s[i]-'a']++
		windowCnt[s[i-pLen]-'a']--
		if windowCnt == targetCnt {
			result = append(result, i)
		}
	}
	return result
}
