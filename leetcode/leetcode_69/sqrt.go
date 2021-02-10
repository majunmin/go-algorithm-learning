/**
  @Author: majm@ushareit.com
  @date: 2021/2/9
  @note:
**/
package leetcode_69

import "fmt"

// 二分查找
// https://leetcode-cn.com/problems/sqrtx/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china
func mySqrt(x int) int {
	left, right := 0, x
	//ans := -1
	for left < right {
		//mid := (left + right) >> 1  // 中间靠左位置，这样写有溢出风险 -> mid := left + (right - left) >> 1
		mid := (left + right + 1) >> 1 // 中间靠右位置           -> mid := left + (right - left + 1) >> 1
		res := mid * mid
		fmt.Println("mid", mid)
		if res == x {
			return mid
		}
		if res < x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// 走到这里退出时 最后退出时  left == right
	return left
}
