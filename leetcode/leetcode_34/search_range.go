/**
  @Author: majm@ushareit.com
  @date: 2021/2/10
  @note:
**/
package leetcode_34

// 排序数组中查找 target 出现的 开始和结束位置
func searchRange(nums []int, target int) []int {
	//return byMe(nums, target)
	return rangeSearch(nums, target)
}

// 二分查找左右逼近的方式
func rangeSearch(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	lIdx := searchLeft(nums, target)
	rIdx := searchRight(nums, target)
	if nums[lIdx] != target || lIdx > rIdx {
		return res
	}
	return []int{lIdx, rIdx}
}

func searchRight(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}

func searchLeft(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func byMe(nums []int, target int) []int {
	res := []int{-1, -1}
	length := len(nums)
	if length == 0 {
		return res
	}
	left, right := 0, length-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return res
	}
	res[0] = left
	end := left
	for i := left; i <= length-1; i++ {
		if nums[i] == target {
			end = i
			continue
		}
		break
	}
	res[1] = end
	return res
}
