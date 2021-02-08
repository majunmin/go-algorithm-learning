/**
  @Author: majm@ushareit.com
  @date: 2021/2/8
  @note:
**/
package sort

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		curVal := arr[i]
		preIdx := i - 1
		for preIdx >= 0 && arr[preIdx] > curVal {
			arr[preIdx+1] = arr[preIdx]
			preIdx--
		}
		arr[preIdx+1] = curVal
	}
}
