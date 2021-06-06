/**
* Date: 2021/6/6 9:01 上午
* Author: majunmin
**/
package leetcode_0039

import "sort"

// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {

	return solution2(candidates, target)
}

func solution2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrace(&result, candidates, 0, []int{}, target)
	return result
}

// 解法二 比较巧妙
func backtrace(result *[][]int, candidates []int, begin int, path []int, target int) {
	// terminate
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := begin; i < len(candidates); i++ {
		num := candidates[i]
		if target-num < 0 {
			break
		}
		path = append(path, num)
		backtrace(result, candidates, i, path, target-num)
		path = path[:len(path)-1]
	}
}

// 时间复杂度
// 空间复杂度
func solution1(candidates []int, target int) [][]int {
	// for 剪枝
	sort.Ints(candidates)
	var result [][]int
	backTtrace(&result, candidates, []int{}, 0, 0, target)
	return result
}

func backTtrace(result *[][]int, candidates []int, path []int, level, sum, target int) {
	// terminate

	for i := level; i < len(candidates); i++ {
		num := candidates[i]
		sum += num
		if sum >= target {
			// save result
			if sum == target {
				path = append(path, num)
				temp := make([]int, len(path))
				copy(temp, path)
				*result = append(*result, temp)
				sum -= num
				path = path[:len(path)-1]
			}
			sum -= num
			return
		}

		path = append(path, num)
		backTtrace(result, candidates, path, i, sum, target)
		// revert
		sum -= num
		path = path[:len(path)-1]
	}
}
