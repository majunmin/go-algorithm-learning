/**
* Date: 2021/6/6 8:33 上午
* Author: majunmin
**/
package leetcode_0047

// 全排列 2
// https://leetcode-cn.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	dfs(&result, nums, []int{}, 0, used)
	return result

}

// 画图解决
//     同一层级不允许出现相同数字
//     同一路径不能有相同选择
//时间复杂度
//空间复杂度
func dfs(result *[][]int, nums, path []int, level int, used []bool) {
	if level == len(nums) {
		temp := make([]int, level)
		copy(temp, path)
		*result = append(*result, temp)
	}

	visit := make(map[int]struct{})
	// for choice in choiceList
	for i, num := range nums {
		// 剪枝
		if used[i] {
			continue
		}
		if _, exist := visit[num]; exist {
			continue
		}

		used[i] = true
		// 记录本层 使用过的数字
		visit[num] = struct{}{}
		path = append(path, num)
		dfs(result, nums, path, level+1, used)

		// revert
		used[i] = false
		path = path[:len(path)-1]
	}
}
