/**
  @Author: majm@ushareit.com
  @date: 2021/1/23
  @note:
**/
package leetcode_79

type pair struct {
	x int
	y int
}

var directions = []pair{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func exist(board [][]byte, word string) bool {

	h, w := len(board), len(board[0])
	// 标记已访问过得
	visits := make([][]bool, h)
	for i := 0; i < len(visits); i++ {
		visits[i] = make([]bool, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if dfs(board, i, j, word, 0, visits) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, word string, k int, visits [][]bool) bool {
	// terminated condition
	if board[i][j] != word[k] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	// process
	h, w := len(board), len(board[0])

	visits[i][j] = true
	var result bool
	for _, direction := range directions {
		if newI, newJ := i+direction.x, j+direction.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !visits[newI][newJ] {
			// 递归
			if dfs(board, newI, newJ, word, k+1, visits) {
				result = true
				break
			}
		}
	}
	// revert state
	visits[i][j] = false
	return result
}
