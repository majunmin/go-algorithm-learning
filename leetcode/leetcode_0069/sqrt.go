/**
  @Author: majm@ushareit.com
  @date: 2021/2/9
  @note:
**/
package leetcode_0069

// 二分查找
// https://leetcode-cn.com/problems/sqrtx/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china
//func mySqrt(x int) int {
//	left, right := 0, x
//	//ans := -1
//	for left < right {
//		//mid := (left + right) >> 1  // 中间靠左位置，这样写有溢出风险 -> mid := left + (right - left) >> 1
//		mid := (left + right + 1) >> 1 // 中间靠右位置           -> mid := left + (right - left + 1) >> 1
//		res := mid * mid
//		fmt.Println("mid", mid)
//		if res == x {
//			return mid
//		}
//		if res < x {
//			left = mid
//		} else {
//			right = mid - 1
//		}
//	}
//	// 走到这里退出时 最后退出时  left == right
//	return left
//}

// https://leetcode-cn.com/problems/sqrtx/
// 时间复杂度 O(logn)
// 空间复杂度 O(1)
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x

	// 退出循环时, left == right 成立
	for left < right {
		// 这里为什么这样分？
		// 1. mid = (left + right) >> 1,    mid 偏左,
		// 1. mid = (left + right + 1) >> 1,mid 偏右
		mid := (left + right + 1) >> 1

		if mid > x/mid { //避免乘法溢出,改用除法
			right = mid - 1
		} else {
			left = mid
		}
	}
	// 走到这里退出时  left == right
	return left
}
