/**
  @Author: majm@ushareit.com
  @date: 2021/3/13
  @note:
**/
package leetcode_0056

// 合并区间
func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length < 2 {
		return intervals
	}
	QuickSort(&intervals)

	var result [][]int
	for i := 1; i < length; i++ {
		if intervals[i-1][1] >= intervals[i][0] {
			if intervals[i-1][1] > intervals[i][1] {
				intervals[i][1] = intervals[i-1][1]
			}
			intervals[i][0] = intervals[i-1][0]
		} else {
			result = append(result, intervals[i-1])
		}
	}
	result = append(result, intervals[length-1])
	return result
}

func QuickSort(intervals *[][]int) {
	left, right := 0, len(*intervals)-1
	quickSortHelper(intervals, left, right)
}

func quickSortHelper(intervals *[][]int, left int, right int) {
	if left >= right {
		return
	}
	partIdx := partition(intervals, left, right)

	quickSortHelper(intervals, left, partIdx-1)
	quickSortHelper(intervals, partIdx+1, right)
}

func partition(intervals *[][]int, left int, right int) int {
	pivot, counter := right, left
	for i := left; i < right; i++ {
		if (*intervals)[i][0] < (*intervals)[pivot][0] {
			(*intervals)[counter], (*intervals)[i] = (*intervals)[i], (*intervals)[counter]
			counter++
		}
	}
	(*intervals)[pivot], (*intervals)[counter] = (*intervals)[counter], (*intervals)[pivot]
	return counter
}
