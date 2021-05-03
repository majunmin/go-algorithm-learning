/**
  @Author: majm@ushareit.com
  @date: 2021/3/10
  @note:
**/
package leetcode_0046

func permute(nums []int) [][]int {
	var result [][]int
	permuteHelper(nums, []int{}, 0, &result)
	return result
}

func permuteHelper(nums, curPath []int, length int, result *[][]int) {
	if len(nums) == length {
		*result = append(*result, curPath)
		return
	}

	// for choice in choice list
	for _, num := range nums {
		if arrContains(curPath, num) {
			continue
		}
		curPath = append(curPath, num)
		permuteHelper(nums, curPath, length+1, result)
		// restore state
		curPath = curPath[:length]
	}
}

// 相当于一个查找算法
func arrContains(arr []int, num int) bool {
	for _, e := range arr {
		if e == num {
			return true
		}
	}
	return false

	//if len(arr) == 0 {
	//	return false
	//}
	//tmp := make([]int, len(arr))
	//copy(tmp, arr)
	//sor.Ints(tmp)
	//idx := sor.SearchInts(tmp, num)
	//return idx < len(arr) && tmp[idx] == num
}
