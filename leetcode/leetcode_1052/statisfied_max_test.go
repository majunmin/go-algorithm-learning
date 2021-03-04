/**
  @Author: majm@ushareit.com
  @date: 2021/2/23
  @note:
**/
package leetcode_1052

import (
	"fmt"
	"testing"
)

func TestMaxSatistified(t *testing.T) {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	fmt.Println(maxSatisfied(customers, grumpy, 3))
}
