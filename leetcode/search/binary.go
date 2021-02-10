/**
  @Author: majm@ushareit.com
  @date: 2021/2/9
  @note:
**/
package search

import (
	"fmt"
)

// arr 是一个有序数组
// 找到 target 在 arr中的索引位置,如果没有那么返回 -1
func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1
	// 如果 len(arr) 是奇数， mid 是中间位置
	//                 偶数, mid 是中间偏左位置

	for left <= right {
		//mid := left + (right-left)>>1  这样写是为了 防止  left + right 溢出
		mid := (left + right) >> 1
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		fmt.Println("mid : ", mid)
	}

	fmt.Println(left)
	fmt.Println(right)

	return -1
}
