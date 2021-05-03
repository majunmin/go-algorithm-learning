/**
  @Author: majm@ushareit.com
  @date: 2021/3/9
  @note:
**/
package exercise

import (
	"fmt"
	"testing"
)

func TestGenParthen(t *testing.T) {
	res := generateParenthesis(3)
	fmt.Println(res)

	sort := QuickSort([]int{2, 5, 7, 4, 99, 0, -3})
	fmt.Println(sort)
}
