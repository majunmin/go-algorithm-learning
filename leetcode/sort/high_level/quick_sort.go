/**
  @Author: majm@ushareit.com
  @date: 2021/2/8
  @note:
**/
package high_level

func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	partIdx := partition(arr, start, end)
	quickSort(arr, start, partIdx-1)
	quickSort(arr, partIdx+1, end)
}

// 找到 arr start 到 end 区间 的一个中间索引,左边的值都比 arr[cnt] 的值小， 右边的都比 arr[cnt] 的值大
func partition(arr []int, start int, end int) int {
	// 选中一个标的物
	pivot := end
	cnt := start

	for i := start; i < end; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[cnt] = arr[cnt], arr[i]
			cnt++
		}
	}
	arr[pivot], arr[cnt] = arr[cnt], arr[pivot]
	return cnt
}

//func partition(arr []int, start int, end int) int {
//	// 选中一个标的物
//	pivot,idx := start,start+1
//	for i := idx; i <= end; i++ {
//		if arr[i] < arr[pivot] {
//			arr[i], arr[idx] = arr[idx], arr[i]
//			idx++
//		}
//	}
//	arr[pivot], arr[idx-1] = arr[idx-1], arr[pivot]
//	return idx-1
//}
