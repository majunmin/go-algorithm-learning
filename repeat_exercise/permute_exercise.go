/**
* Date: 2021/6/4 7:52 上午
* Author: majunmin
**/
package repeat_exercise

// https://leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int {
	var result [][]int
	//dfs(&result, nums, len(nums), make([]int, 0, 4))

	used := make([]bool, 4)
	dfs(&result, nums, len(nums), []int{}, used)
	return result

}

// used 记录已经选过的 choice
func dfs(result *[][]int, nums []int, length int, path []int, used []bool) {
	if length == len(path) {
		//由于 切片是引用类型,索引 这里 要将拷贝的结果 放进  result中
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// for choice in choiceList
	for i, num := range nums {
		if used[i] {
			continue
		}
		used[i] = true
		newUsed := make([]bool, length)
		copy(newUsed, used)

		// drill down
		path = append(path, num)
		dfs(result, nums, length, path, newUsed)
		// revert
		path = path[:len(path)-1]
		used[i] = false
	}
}
