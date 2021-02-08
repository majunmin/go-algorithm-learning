/**
  @Author: majm@ushareit.com
  @date: 2021/2/8
  @note:
**/
package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 9, 1, 3, 7, 10, 8}
	InsertionSort(arr)
	fmt.Println(arr)
}
