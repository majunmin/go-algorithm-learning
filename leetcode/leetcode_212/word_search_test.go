/**
  @Author: majm@ushareit.com
  @date: 2021/1/24
  @note:
**/
package leetcode_212

import (
	"fmt"
	"testing"
)

func TestWordSearch(t *testing.T) {
	var board = [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	words := []string{"oath", "pea", "eat", "rain"}

	fmt.Println(findWords(board, words))
}
