/**
  @Author: majm@ushareit.com
  @date: 2021/2/9
  @note:
**/
package leetcode_278

var ver = 4

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right := 0, n
	for left < right {
		mid := (left + right + 1) / 2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return version >= ver
}
