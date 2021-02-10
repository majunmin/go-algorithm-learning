/**
  @Author: majm@ushareit.com
  @date: 2021/2/10
  @note:
**/
package leetcode_154

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
// 这个题目官方题解讲的还是挺清晰的 --  对二分查找的 边界条件的练习
func findMin(nums []int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	return nums[left]
}
