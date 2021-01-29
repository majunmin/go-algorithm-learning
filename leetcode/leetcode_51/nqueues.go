/**
  @Author: majm@ushareit.com
  @date: 2021/1/28
  @note:
**/
package leetcode_51

func solveNQueens(n int) [][]string {
	//solveNQueens1(n)
	return solveNQueens2(n)
}

func solveNQueens2(n int) [][]string {
	if n <= 0 {
		return nil
	}
	var result [][]string

	// init board
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = '.'
		}
	}
	//  矩阵特性 同一 col上   col = constant
	//  矩阵特性 同一 na 上   col - row = constant
	//  矩阵特性 同一 pie上   col+row = constant
	// 用于判断合法性  set
	colSet := make(map[int]struct{})
	pieSet := make(map[int]struct{})
	naSet := make(map[int]struct{})

	backtrace2(board, 0, colSet, pieSet, naSet, &result)
	return result

}

func backtrace2(board [][]byte, curIdx int, colSet map[int]struct{}, pieSet map[int]struct{}, naSet map[int]struct{}, result *[][]string) {

	n := len(board)
	if curIdx == len(board) {
		saveInRes(board, result)
		return
	}

	whiteRow := generateRow(n)
	for i := 0; i < n; i++ {
		if isValid(board, curIdx, i, colSet, pieSet, naSet) {
			whiteRow[i] = 'Q'
			board[curIdx] = whiteRow
			colSet[i] = struct{}{}
			pieSet[curIdx+i] = struct{}{}
			naSet[curIdx-i] = struct{}{}

			backtrace(board, curIdx+1, result)
			// revert
			whiteRow[i] = '.'
			delete(colSet, i)
			delete(pieSet, curIdx+i)
			delete(naSet, curIdx-i)
		}
	}

}

func isValid(board [][]byte, row int, col int, colSet map[int]struct{}, pieSet map[int]struct{}, naSet map[int]struct{}) bool {
	_, ok1 := colSet[col]
	if ok1 {
		return false
	}
	_, ok2 := pieSet[row+col]
	if ok2 {
		return false
	}
	_, ok3 := naSet[row-col]
	if ok3 {
		return false
	}
	return true
}

func solveNQueens1(n int) [][]string {
	var result [][]string
	board := make([][]byte, n)
	// init board
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	backtrace(board, 0, &result)

	return result
}

// 回溯算法
func backtrace(board [][]byte, curIdx int, result *[][]string) {
	if curIdx == len(board) {
		saveInRes(board, result)
		return
	}

	curRow := generateRow(len(board[0]))

	// for choice in choiceList
	//     do choice
	//     backtrace
	//     revert choice
	for i := 0; i < len(board[0]); i++ {
		curRow[i] = 'Q'
		if validate(curIdx, i, board) {
			board[curIdx] = curRow
			backtrace(board, curIdx+1, result)
		}
		// revert
		curRow[i] = '.'
	}
}

func generateRow(len int) []byte {
	curRow := make([]byte, len)
	for i := 0; i < len; i++ {
		curRow[i] = '.'
	}
	return curRow
}

func validate(row int, col int, board [][]byte) bool {
	// validate col
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// validate pie
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// validate na
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func saveInRes(board [][]byte, result *[][]string) {
	res := make([]string, 0)
	for _, row := range board {
		res = append(res, string(row))
	}
	*result = append(*result, res)
}

func generateRows(n int) [][]byte {
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if i == j {
				row[j] = 'Q'
			}
			row[j] = '.'
		}
		res = append(res, row)
	}
	return res
}
