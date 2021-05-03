/**
  @Author: majm@ushareit.com
  @date: 2021/3/4
  @note:
**/
package leetcode_0300

// https://leetcode-cn.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	l := len(nums)
	lenRes := make([]int, l)
	lenRes[0] = 1

	res := 0
	for i := 1; i < l; i++ {
		lenRes[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				lenRes[i] = max(lenRes[j]+1, lenRes[i])
			}
		}
		res = max(res, lenRes[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
