/**
  @Author: majm@ushareit.com
  @date: 2021/1/29
  @note:
**/
package leetcode_37

import (
	"go-algorithm-learning/leetcode/leetcode_36"
)

// 解数独
// https://leetcode-cn.com/problems/sudoku-solver/
func solveSudoku(board [][]byte) {
	//solve(&board)
	for i := 0; i < 72; i++ {
		solve(&board)
	}

}

// 逐个位置遍历  row = num/9  col  = num %9
func solve(board *[][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*board)[i][j] != '.' {
				continue
			}
			var c byte
			for c = '1'; c <= '9'; c++ {
				(*board)[i][j] = c
				if leetcode_36.IsValidSudoku(*board) {
					if solve(board) {
						return true
					}
				}
			}
			// 该点放置 1-9 都不可以
			(*board)[i][j] = '.'
			return false
		}
	}
	return true
}
