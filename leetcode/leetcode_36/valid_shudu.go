/**
  @Author: majm@ushareit.com
  @date: 2021/1/29
  @note:
**/
package leetcode_36

// 验证有效地 数独
// 格式为 9*9
func IsValidSudoku(board [][]byte) bool {

	return solution2(board)

}

func solution2(board [][]byte) bool {
	// init byte
	rows := make([][]byte, 9)
	cols := make([][]byte, 9)
	boxes := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]byte, 9)
		cols[i] = make([]byte, 9)
		boxes[i] = make([]byte, 9)
		//for j := 0; j < 9; j++ {
		//	rows[i][j] = '0'
		//	cols[i][j] = '0'
		//	boxes[i][j] = '0'
		//}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			if rows[i][num-'1'] == 1 {
				return false
			}
			rows[i][num-'1'] = 1

			if cols[j][num-'1'] == 1 {
				return false
			}
			cols[j][num-'1'] = 1

			if boxes[(i/3)*3+j/3][num-'1'] == 1 {
				return false
			}
			boxes[i/3*3+j/3][num-'1'] = 1
		}
	}
	return true
}

// hiahia
func lowSolution(board [][]byte) bool {
	choices := [][]byte{
		{'0', '1', '2'},
		{'3', '4', '5'},
		{'6', '7', '8'},
	}

	for i := 0; i < 9; i++ {
		colSet := make(map[byte]struct{})
		rowSet := make(map[byte]struct{})

		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, ok := rowSet[board[i][j]]; ok {
					return false
				}
				rowSet[board[i][j]] = struct{}{}
			}

			if board[j][i] != '.' {
				if _, ok := colSet[board[j][i]]; ok {
					return false
				}
				colSet[board[j][i]] = struct{}{}
			}
		}
	}

	for _, rows := range choices {
		for _, cols := range choices {
			blockSet := make(map[byte]struct{})
			for _, i := range rows {
				for _, j := range cols {
					if board[i][j] != '.' {
						if _, ok := blockSet[board[i][j]]; ok {
							return false
						}
						blockSet[board[i][j]] = struct{}{}
					}
				}
			}
		}
	}
	return true
}
