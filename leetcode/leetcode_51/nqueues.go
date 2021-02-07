/**
  @Author: majm@ushareit.com
  @date: 2021/1/28
  @note:
**/
package leetcode_51

// 有个公式需要记住  方便分析 位运算 和 后面的 set 解法
//  na = n - 1 - row + col
//  shu = col
//  pie = row + col
// https://zhuanlan.zhihu.com/p/22846106
func solveNQueens(n int) [][]string {
	//return solveNQueens1(n)
	//return solveNQueens2(n)
	//return solveNQueens3(n)
	return bitSolution1(n)
}

// shu 冲突位置 shu
// pie 冲突位置 pie >> row
// na 冲突位置  na  >> n-1-row
func bitSolution1(n int) [][]string {
	result := make([][]string, 0)
	board := make([][]byte, n)
	shu, pie, na := 0, 0, 0
	dfs(0, n, shu, pie, na, board, &result)
	return result
}

// 用 位运算 代替  set
func dfs(level int, n int, shu, pie, na int, board [][]byte, result *[][]string) {
	if level == n {
		saveInRes(board, result)
		return
	}

	whiteRow := generateRow(n)
	for col := 0; col < n; col++ {
		j := level + col
		k := n - 1 - level + col

		// 表示该位已经被占用
		if (shu>>col|pie>>j|na>>k)&1 == 1 {
			continue
		}
		shu |= 1 << col
		pie |= 1 << (level + col)
		na |= 1 << (n - 1 - level + col)
		whiteRow[col] = 'Q'
		board[level] = whiteRow

		dfs(level+1, n, shu, pie, na, board, result)

		whiteRow[col] = '.'
		board[level] = nil
		shu &= ^(1 << col)
		pie &= ^(1 << (level + col))
		na &= ^(1 << (n - 1 - level + col))
	}

}

var solutions [][]string

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
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
