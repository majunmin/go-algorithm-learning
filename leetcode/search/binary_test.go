/**
  @Author: majm@ushareit.com
  @date: 2021/2/9
  @note:
**/
package search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	idx := BinarySearch(arr, 9)
	fmt.Println(idx)
}
