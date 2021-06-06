/**
* Date: 2021/6/3 1:02 上午
* Author: majunmin
**/
package leetcode_0077

// https://leetcode-cn.com/problems/combinations/
func combine(n int, k int) [][]int {
	if n < k || k <= 0 {
		return nil
	}
	var result [][]int

	dfs(&result, []int{}, 1, n, k)
	return result
}

func dfs(result *[][]int, path []int, level int, n int, limit int) {
	if len(path) == limit {
		temp := make([]int, limit)
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// for choice in choiceList
	for i := level; i <= n-limit+len(path)+1; i++ {
		path = append(path, i)
		// 开启下一轮搜索 ,搜索起点+ 1,因为 组合 不允许出现重复
		dfs(result, path, i+1, n, limit)
		// revert
		path = path[:len(path)-1]
	}
}
