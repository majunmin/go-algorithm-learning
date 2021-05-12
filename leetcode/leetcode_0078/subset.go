/**
* Date: 2021/5/11 8:15 上午
* Author: majunmin
**/
package leetcode_0078

// https://leetcode-cn.com/problems/subsets/
func subsets(nums []int) [][]int {
	return iterSolution(nums)
}

// 000
// 001
// 010
// 011
// 100
// 101
// 110
// 111
// 位运算
func iterSolution(nums []int) [][]int {
	var result [][]int
	for i := 0; i < 1<<len(nums); i++ {
		path := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i&(1<<j) != 0 {
				path = append(path, nums[j])
			}
		}
		result = append(result, path)
	}
	return result
}

// 回溯算法
func huisu1(nums []int) [][]int {
	var result [][]int
	backTrace(&result, nums, []int{}, 0)
	return result
}

// 回溯
func backTrace(result *[][]int, nums []int, path []int, level int) {
	// 此处不能将  path  直接放入 result, 切片是引用类型
	dstPath := make([]int, len(path))
	copy(dstPath, path)
	*result = append(*result, dstPath)
	if level == len(nums) {
		return
	}

	// for choice in choiceList
	for i := level; i < len(nums); i++ {
		path = append(path, nums[i])
		backTrace(result, nums, path, i+1)
		path = path[:len(path)-1]
	}
}
