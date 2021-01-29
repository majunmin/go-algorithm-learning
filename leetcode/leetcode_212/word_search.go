/**
  @Author: majm@ushareit.com
  @date: 2021/1/23
  @note:
**/
package leetcode_212

import (
	"go-algorithm-learning/leetcode/leetcode_208"
)

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func findWords(board [][]byte, words []string) []string {

	trie := leetcode_208.Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	h, w := len(board), len(board[0])

	// init visits(flag the byte has visit)
	visits := make([][]bool, h)
	for i := 0; i < h; i++ {
		visits[i] = make([]bool, w)
	}

	result := make(map[string]struct{}, len(words))

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			dfs(board, i, j, string(board[i][j]), visits, trie, &result)
		}
	}
	var res []string
	for w := range result {
		res = append(res, w)
	}
	return res
}

func dfs(board [][]byte, i int, j int, curWord string, visits [][]bool, trie leetcode_208.Trie, result *map[string]struct{}) {
	// terminate
	if !trie.StartsWith(curWord) {
		return
	}

	if trie.Search(curWord) {
		(*result)[curWord] = struct{}{}
	}

	visits[i][j] = true
	h, w := len(board), len(board[0])

	for _, direction := range directions {
		if newI, newJ := i+direction[0], j+direction[1]; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !visits[newI][newJ] {
			dfs(board, newI, newJ, curWord+string(board[newI][newJ]), visits, trie, result)
		}
	}
	// revert stat
	visits[i][j] = false
}
