/**
  @Author: majm@ushareit.com
  @date: 2021/1/23
  @note:
**/
package leetcode_79

import (
	"testing"
)

func TestWordSearch(t *testing.T) {

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	println(exist(board, "ABCCED"))
}
