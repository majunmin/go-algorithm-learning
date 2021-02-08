/**
  @Author: majm@ushareit.com
  @date: 2021/2/8
  @note:
**/
package high_level

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{3, 7, 1, 2, 9, 11, 8}
	QuickSort(arr)
	fmt.Println(arr)
}
