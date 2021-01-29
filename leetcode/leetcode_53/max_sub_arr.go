/**
  @Author: majm@ushareit.com
  @date: 2021/1/17
  @note:
**/
package leetcode_53

import "go-algorithm-learning/leetcode/common"

func maxSubArray(nums []int) int {
	return dpMaxSubArr(nums)
}

/**
贪婪算法

如果之前和 <0, 丢弃之前和 当前和 为  当前值
max = common.MaxInt(curSum, max)
*/
func greedySubArr(nums []int) int {
	curSum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = common.MaxInt(nums[i], curSum+nums[i])
		max = common.MaxInt(curSum, max)
	}
	return max
}

func dpMaxSubArr(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = common.MaxInt(nums[i], nums[i]+nums[i-1])
		max = common.MaxInt(max, nums[i])
	}
	return max
}
