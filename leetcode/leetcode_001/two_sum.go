/**
  @Author: majm@ushareit.com
  @date: 2020/12/2
  @note:
**/
package leetcode_001

import (
	"sort"
)

// [两数之和](https://leetcode-cn.com/problems/two-sum/)
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	return Solution1(nums, target)
}

func Solution1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; i < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 因为要返回下标，这种方法要记下下标
func Solution2(nums []int, target int) []int {
	sort.Ints(nums)
	length := len(nums)
	left := 0
	right := length - 1
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{left, right}
		}
		if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// hash表法
func Solution3(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, num := range nums {
		if p, ok := hashMap[target-num]; ok {
			return []int{p, i}
		}
		hashMap[num] = i
	}
	return nil
}
