/**
* Date: 2021/5/3 3:22 下午
* Author: majunmin
**/
package exercise

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 1, 9, 3, 0, 4, -1, 6}
	//InsertSort(arr)
	QuickSort(arr)
	fmt.Println(arr)

	//sort := QuickSort([]int{2, 5, 7, 4, 99, 0, -3})
	//fmt.Println(sort)
}
