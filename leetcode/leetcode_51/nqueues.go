/**
  @Author: majm@ushareit.com
  @date: 2021/1/28
  @note:
**/
package leetcode_51

func solveNQueens(n int) [][]string {
	//return solveNQueens1(n)
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
		if isValid(curIdx, i, colSet, pieSet, naSet) {
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

func isValid(row int, col int, colSet map[int]struct{}, pieSet map[int]struct{}, naSet map[int]struct{}) bool {
	if _, ok := colSet[col]; ok {
		return false
	}

	if _, ok := pieSet[row+col]; ok {
		return false
	}

	if _, ok := naSet[row-col]; ok {
		return false
	}
	return true
}

func solveNQueens1(n int) [][]string {
	var result [][]string
	board := make([][]byte, n)
	backtrace(board, 0, &result)

	return result
}

// 回溯算法
func backtrace(board [][]byte, level int, result *[][]string) {
	n := len(board)
	if level == n {
		saveInRes(board, result)
		return
	}

	whiteRow := generateRow(n)

	// for choice in choiceList
	//     do choice
	//     backtrace
	//     revert choice
	for i := 0; i < n; i++ {
		whiteRow[i] = 'Q'
		if validate(level, i, board) {
			board[level] = whiteRow
			backtrace(board, level+1, result)
		}
		// revert
		whiteRow[i] = '.'
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
	if row == 0 {
		return true
	}
	// validate col
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// validate pie
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// validate na
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
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
