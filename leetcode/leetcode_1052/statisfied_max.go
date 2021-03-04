/**
  @Author: majm@ushareit.com
  @date: 2021/2/23
  @note:
**/
package leetcode_1052

// https://leetcode-cn.com/problems/grumpy-bookstore-owner/
// 时间复杂度  O(N)
// 空间复杂度  O(1)
func maxSatisfied(customers []int, grumpy []int, X int) int {
	total := 0
	length := len(grumpy)
	// 先计算出  一定有的 顾客数
	for i := 0; i < length; i++ {
		total += (1 - grumpy[i]) * customers[i]
	}

	increase := 0
	for i := 0; i < X; i++ {
		increase += grumpy[i] * customers[i]
	}
	maxIncrease := increase

	for i := X; i < length; i++ {
		increase = increase - grumpy[i-X]*customers[i-X] + grumpy[i]*customers[i]
		if increase > maxIncrease {
			maxIncrease = increase
		}
	}
	return total + maxIncrease
}
