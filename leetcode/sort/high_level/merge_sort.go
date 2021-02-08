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
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	// 如果 right - left + 1 = 偶数,mid取中间偏左位置，所以下面 mergeSort 取 (mid + 1) ~ right
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

// 合并两个有序数组 需要用到额外的空间
func merge(arr []int, left int, mid int, right int) {
	tmpArr := make([]int, right-left+1)
	idx := 0
	i1, i2 := left, mid+1
	for i1 <= mid && i2 <= right {
		if arr[i1] < arr[i2] {
			tmpArr[idx] = arr[i1]
			idx++
			i1++
		} else {
			tmpArr[idx] = arr[i2]
			idx++
			i2++
		}
	}

	for i1 <= mid {
		tmpArr[idx] = arr[i1]
		idx++
		i1++
	}
	for i2 <= right {
		tmpArr[idx] = arr[i2]
		idx++
		i2++
	}

	for i := 0; i < right-left+1; i++ {
		arr[left+i] = tmpArr[i]
	}
}
