/**
* Date: 2021/5/12 12:53 上午
* Author: majunmin
**/
package leetcode_0090

import "sort"

// https://leetcode-cn.com/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	backtrace([]int{}, nums, 0, &res)
	return res
}

// 剪枝:  路径有重复的  可以用一个set 剔除重复的(processedNum) 去重
func backtrace(path []int, nums []int, level int, res *[][]int) {
	dst := make([]int, len(path))
	copy(dst, path)
	*res = append(*res, dst)
	length := len(nums)
	// terminator condition
	if level == length {
		return
	}

	// for choice in choiceListr
	for i := level; i < length; i++ {
		if i > level && nums[i] == nums[i-1] {
			continue
		}
		backtrace(append(path, nums[i]), nums, i+1, res)
		// revert
	}
}
