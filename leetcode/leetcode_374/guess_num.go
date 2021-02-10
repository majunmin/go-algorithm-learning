/**
  @Author: majm@ushareit.com
  @date: 2021/2/9
  @note:
**/
package leetcode_374

/**
-- 题目看了半天...
https://leetcode-cn.com/problems/guess-number-higher-or-lower/

 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
*/
// 时间复杂度  O(lgN)
// 空间复杂度  O(1)  没有使用额外的空间
var pick = 2

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) >> 1
		res := guess(mid)
		if res == 0 {
			return mid
		}
		if res > 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func guess(target int) int {
	if pick > target {
		return 1
	} else if pick < target {
		return -1
	}
	return 0
}
