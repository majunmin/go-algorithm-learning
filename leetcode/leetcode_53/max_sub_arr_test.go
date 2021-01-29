/**
  @Author: majm@ushareit.com
  @date: 2021/1/18
  @note:
**/
package leetcode_53

import (
	"fmt"
	"testing"
)

func TestMaxSubArr(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
