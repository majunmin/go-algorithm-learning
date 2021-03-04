/**
  @Author: majm@ushareit.com
  @date: 2021/3/2
  @note:
**/
package leetcode_0304

// https://leetcode-cn.com/problems/range-sum-query-2d-immutable/
type NumMatrix struct {
	preSumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	rows, cols := len(matrix), len(matrix[0])
	if cols == 0 {
		return NumMatrix{}
	}
	preSumMatrix := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		preSumMatrix[i] = make([]int, cols+1)
	}
	// init preSumMatrix
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			preSumMatrix[i][j] = preSumMatrix[i][j-1] + preSumMatrix[i-1][j] - preSumMatrix[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{preSumMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	matrix := this.preSumMatrix
	if matrix == nil {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	if row1 < 0 || row1 > rows || row2 < 0 || row2 > rows || col1 < 0 || col1 > cols || col2 < 0 || col2 > cols {
		return 0
	}
	return matrix[row2+1][col2+1] - matrix[row2+1][col1] - matrix[row1][col2+1] + matrix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
