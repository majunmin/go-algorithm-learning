/**
  @Author: majm@ushareit.com
  @date: 2021/3/10
  @note:
**/
package leetcode_0015

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	//sum := threeSum([]int{-2,0,0,2,2})
	sum := threeSum2([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0})
	sum2 := threeSum2([]int{1, -1, -1, 0})
	fmt.Println(sum)
	fmt.Println(sum2)
}
