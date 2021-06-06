/**
  @Author: majm@ushareit.com
  @date: 2021/3/9
  @note:
**/
package leetcode_0015

import (
	"sort"
)

// -1 0 1 2 -1  -4
// -4 -1 -1 0 1 2
func threeSum(nums []int) [][]int {
	length := len(nums)
	if length <= 2 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0)
	for i := 0; i <= length-3; i++ {
		// 去重
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, length-1
		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func threeSum2(nums []int) [][]int {

	var result [][]int
	length := len(nums)
	if length < 3 {
		return result
	}

	// sort
	sort.Ints(nums)

	for i := 0; i < length; i++ {
		num := nums[i]
		if num > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 判断 双指针判断  两数之和
		l, r := i+1, length-1
		for l < r {
			if l < r && nums[l]+nums[r]+num > 0 {
				r--
			}
			if l < r && nums[l]+nums[r]+num < 0 {
				l++
			}
			if l < r && nums[l]+nums[r]+num == 0 {
				result = append(result, []int{num, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return result
}
