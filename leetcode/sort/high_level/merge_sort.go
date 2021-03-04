/**
  @Author: majm@ushareit.com
  @date: 2021/2/8
  @note:
**/
package high_level

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left int, right int) {
	if right <= left {
		return
	}

	mid := (left + right) >> 1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

// left -> mid 是有序的
// mid + 1 -> right 是有序的
// merge的作用就是 合并两个有序数组
func merge(arr []int, left int, mid int, right int) {
	tempArr := make([]int, right-left+1)
	idx := 0
	lIdx := left
	rIdx := mid + 1
	for lIdx <= mid && rIdx <= right {
		if arr[lIdx] < arr[rIdx] {
			tempArr[idx] = arr[lIdx]
			lIdx++
		} else {
			tempArr[idx] = arr[rIdx]
			rIdx++

		}
		idx++
	}

	for lIdx <= mid {
		tempArr[idx] = arr[lIdx]
		lIdx++
		idx++
	}

	for rIdx <= right {
		tempArr[idx] = arr[rIdx]
		rIdx++
		idx++
	}

	for i := 0; i < len(tempArr); i++ {
		arr[left+i] = tempArr[i]
	}
}
