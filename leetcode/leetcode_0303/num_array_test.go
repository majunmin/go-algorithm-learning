/**
  @Author: majm@ushareit.com
  @date: 2021/3/1
  @note:
**/
package leetcode_0303

import "testing"

func TestSumArr(t *testing.T) {
	arr := []int{-2, 0, 3, -5, 2, -1}
	numsArr := Constructor(arr)
	numsArr.SumRange(0, 2)
}
