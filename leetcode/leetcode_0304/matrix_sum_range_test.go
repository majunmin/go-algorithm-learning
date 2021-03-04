/**
  @Author: majm@ushareit.com
  @date: 2021/3/2
  @note:
**/
package leetcode_0304

import "testing"

func TestPreSumMatrix(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := Constructor(matrix)
	numMatrix2 := Constructor(nil)
	numMatrix3 := Constructor([][]int{})
	numMatrix2.SumRegion(0, 0, 0, 0)
	numMatrix3.SumRegion(0, 0, 0, 0)
	numMatrix.SumRegion(2, 1, 4, 3)
	numMatrix.SumRegion(1, 1, 2, 2)
	numMatrix.SumRegion(1, 2, 2, 4)
}
