/**
  @Author: majm@ushareit.com
  @date: 2021/2/9
  @note:
**/
package leetcode_153

import "math"

func findMin(nums []int) int {
	return LogN(nums)
}

// 二分查找
//   时间复杂度:O(logN)
//   空间复杂度:O(1)
// if nums[pivot] <= nums[right] 说明 pivot 在右侧
func LogN(nums []int) int {
	length := len(nums)
	if length == 0 {
		return math.MinInt32
	}

	left, right := 0, length-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}

// 遍历 一遍 数组
//   时间复杂度:O(N)
//   空间复杂度:O(1)
func On(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}
