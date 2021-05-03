/**
  @Author: majm@ushareit.com
  @date: 2021/3/9
  @note:
**/
package exercise

// 插入排序
func InsertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		curVal := arr[i]
		prev := i - 1
		for prev >= 0 && arr[prev] > curVal {
			arr[prev+1] = arr[prev]
		}
		arr[prev+1] = curVal
	}
	return arr
}

// 选择排序
func SelectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIdx := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

// 选出一个 中间值,每次将比他小的值  放在其左,比他大的值放在其右
func QuickSort(arr []int) []int {
	quickSortHelper(arr, 0, len(arr)-1)
	return arr
}

func quickSortHelper(arr []int, left int, right int) {
	// terminate
	if left >= right {
		return
	}
	// process
	partitionIdx := partition(arr, left, right)

	// 递归调用 分支
	quickSortHelper(arr, left, partitionIdx-1)
	quickSortHelper(arr, partitionIdx+1, right)
}

func partition(arr []int, left int, right int) int {
	pivot, counter := right, left
	for i := left; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return pivot
}

// merge sort 分治算法
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mergeSortHelper(arr, 0, len(arr)-1)
	return arr
}

func mergeSortHelper(arr []int, left, right int) {
	// terminate
	if left <= right {
		return
	}

	mid := left + (right-left)>>1
	mergeSortHelper(arr, left, mid)
	mergeSortHelper(arr, mid+1, right)

	merge(arr, left, mid, right)
}

// 合并两个有序数组
func merge(arr []int, left int, mid int, right int) []int {
	tmpArr := make([]int, len(arr))
	lIdx, rIdx, tmpIdx := left, mid+1, 0

	for ; lIdx <= mid && rIdx <= right; tmpIdx++ {
		if arr[lIdx] <= arr[rIdx] {
			tmpArr[tmpIdx] = arr[lIdx]
			lIdx++
		} else {
			tmpArr[tmpIdx] = arr[rIdx]
			rIdx++
		}
	}

	for lIdx <= mid {
		tmpArr[tmpIdx] = arr[lIdx]
		tmpIdx++
		lIdx++
	}
	for rIdx <= right {
		tmpArr[tmpIdx] = arr[rIdx]
		tmpIdx++
		rIdx++
	}

	return tmpArr
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
