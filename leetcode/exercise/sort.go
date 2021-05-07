/**
  @Author: majm@ushareit.com
  @date: 2021/3/9
  @note:
**/
package exercise

// 插入排序
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if temp >= arr[j] {
				break
			}
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}

func MergeSort(arr []int) {
	mergeSortHelper(arr, 0, len(arr)-1)
}

// 将  arr[left - right] 排好序
func mergeSortHelper(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)>>1
	mergeSortHelper(arr, 0, mid)
	mergeSortHelper(arr, mid+1, right)
	merge(arr, left, mid, right)
}

// 合并两个有序数组
func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	idx := 0
	lIdx, rIdx := left, mid+1
	for lIdx <= mid && rIdx <= right {
		if arr[lIdx] < arr[rIdx] {
			temp[idx] = arr[lIdx]
			lIdx++
		} else {
			temp[idx] = arr[rIdx]
			rIdx++
		}
		idx++
	}

	for lIdx <= mid {
		temp[idx] = arr[lIdx]
		lIdx++
		idx++
	}

	for rIdx <= right {
		temp[idx] = arr[rIdx]
		rIdx++
		idx++
	}
	i := left
	for _, item := range temp {
		arr[i] = item
		i++
	}
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	partitionIdx := partition(arr, left, right)
	quickSort(arr, left, partitionIdx-1)
	quickSort(arr, partitionIdx+1, right)
}

// 基准值的选取pivot
// 返回一个索引 partIdx, partIdx 左边的 肯定小于  partIdx 右边的
func partition(arr []int, left, right int) int {
	pivot, counter := right, left
	for i := left; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}
	arr[counter], arr[pivot] = arr[pivot], arr[counter]
	return counter
	//pivot := arr[right]
	//for left < right {
	//	for left < right && arr[left] < pivot {
	//		left++
	//	}
	//	if left < right {
	//		arr[right] = arr[left]
	//		right--
	//	}
	//
	//	for left < right && arr[right] > pivot {
	//		right--
	//	}
	//	if left < right {
	//		arr[left] = arr[right]
	//		left++
	//	}
	//}
	//arr[left] = pivot
	//return left
}
