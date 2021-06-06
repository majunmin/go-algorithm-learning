/**
* Date: 2021/5/30 1:32 下午
* Author: majunmin
**/
package repeat_exercise

func minWindow(s string, t string) string {

	need := make(map[byte]int)   // 需要的计数
	window := make(map[byte]int) // 当前窗口计数器

	valid := 0 // 如果 valid == len(t)  表示合法
	for i := range t {
		need[t[i]]++
	}

	l, r := 0, 0
	minLen, length := len(s), len(s)
	res := ""

	for r < length {
		c := s[r]
		r++
		// r指针右移更新窗口
		if _, exist := need[c]; exist {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		// 根据条件  l 指针左移,收缩窗口
		for l < r && valid == len(need) {
			if r-l <= minLen { // 更新结果
				minLen = r - l
				res = s[l:r]
			}
			d := s[l]
			if _, exist := need[d]; exist {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}

			l++
		}
	}
	return res
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
