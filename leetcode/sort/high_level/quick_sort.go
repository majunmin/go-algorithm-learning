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

func quickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	midIdx := partition(arr, left, right)
	quickSort(arr, left, midIdx-1)
	quickSort(arr, midIdx+1, right)
}

//partition 找到 arr left 到 right 区间 的一个中间索引,左边的值都比 arr[cnt] 的值小， 右边的都比 arr[cnt] 的值大
func partition(arr []int, left int, right int) int {

	pivot := right
	cnt := left
	for i := left; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[cnt] = arr[cnt], arr[i]
			cnt++
		}
	}
	arr[cnt], arr[pivot] = arr[pivot], arr[cnt]
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
