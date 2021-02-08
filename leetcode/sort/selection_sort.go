/**
  @Author: majm@ushareit.com
  @date: 2021/2/8
  @note:
**/
package sort

func SelectionSort(arr []int) {
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
}
