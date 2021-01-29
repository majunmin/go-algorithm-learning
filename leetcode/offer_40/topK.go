/**
  @Author: majm@ushareit.com
  @date: 2021/1/16
  @note:
**/
package offer_40

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	//return code(arr, k)
	return quickSort(arr, k)
}

func quickSort(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	start, end := 0, len(arr)-1
	index := partition(arr, start, end)
	for index != k-1 {
		if index > k-1 {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(arr, start, end)
	}
	return arr[:k] //注意：返回的这k个元素不一定是有序的
}

func partition(arr []int, left, right int) int {
	pir := arr[left]
	for left < right {
		for left < right && arr[right] >= pir {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pir {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pir
	return left
}

func code(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[0:k]
}
