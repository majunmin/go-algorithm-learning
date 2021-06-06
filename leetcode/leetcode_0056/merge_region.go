/**
  @Author: majm@ushareit.com
  @date: 2021/3/13
  @note:
**/
package leetcode_0056

func merge1(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)

	var result [][]int
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if pre[1] >= intervals[i][0] && pre[1] <= intervals[i][1] {
			pre[1] = intervals[i][1]
		} else if pre[1] < intervals[i][0] {
			result = append(result, pre)
			pre = intervals[i]
		}
		//if pre[1] > intervals[i][1] {
		//} else if pre[1] >= intervals[i][0] {
		//	pre[1] =  intervals[i][1]
		//} else {
		//	result = append(result, []int{pre[0], pre[1]})
		//	pre = intervals[i]
		//}
	}
	result = append(result, []int{pre[0], pre[1]})
	return result
}

func quickSort(intervals [][]int, left, right int) {
	if left >= right {
		return
	}
	idx := partitionx(intervals, left, right)
	quickSort(intervals, left, idx-1)
	quickSort(intervals, idx+1, right)
	return
}

func partitionx(intervals [][]int, left, right int) int {
	pivot, cnt := right, left
	for i := left; i < right; i++ {
		if intervals[i][0] < intervals[pivot][0] {
			intervals[i], intervals[cnt] = intervals[cnt], intervals[i]
			cnt++
		}
	}
	intervals[cnt], intervals[pivot] = intervals[pivot], intervals[cnt]
	return cnt
}

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
