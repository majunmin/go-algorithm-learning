/**
  @Author: majm@ushareit.com
  @date: 2021/2/20
  @note:
**/
package leetcode_239

import "testing"

func TestSlidingWindow(t *testing.T) {
	param := []int{1, 3, -1, -3, 5, 3, 6, 7}
	maxSlidingWindow(param, 3)
}
