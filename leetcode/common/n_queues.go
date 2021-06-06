/**
* Date: 2021/6/2 1:19 下午
* Author: majunmin
**/
package common

import "strings"

// https://leetcode-cn.com/problems/n-queens/
func solveNQueens(n int) [][]string {

	var result [][]string
	pieSet := make(map[int]struct{})
	naSet := make(map[int]struct{})
	colSet := make(map[int]struct{})

	solveQueues(&result, []string{}, n, pieSet, naSet, colSet, 0)
	return result
}

// 维护 naset pieSet  colSet 状态
func solveQueues(result *[][]string, path []string, n int, pieSet map[int]struct{}, naSet map[int]struct{}, colSet map[int]struct{}, row int) {
	// terminate
	if row == n {
		*result = append(*result, path)
		return
	}

	// for choice in choiceList
	for i := 0; i < n; i++ {
		if !valid(row, i, colSet, pieSet, naSet) {
			continue
		}
		pieSet[row+i] = struct{}{}
		naSet[row-i] = struct{}{}
		colSet[i] = struct{}{}
		currentPath := buildEmptyPath(n, i)
		path = append(path, currentPath)
		solveQueues(result, path, n, pieSet, naSet, colSet, row+1)
		path = path[:len(path)-1]
		delete(pieSet, row+i)
		delete(naSet, row-i)
		delete(colSet, i)
	}
}

func valid(row int, col int, colSet, pieSet, naSet map[int]struct{}) bool {
	if _, exist := colSet[col]; exist {
		return false
	}
	if _, exist := pieSet[row+col]; exist {
		return false
	}
	if _, exist := naSet[row-col]; exist {
		return false
	}
	return true
}

func buildEmptyPath(n, i int) string {
	str := strings.Repeat(".", n)
	return str[:i] + "Q" + str[i+1:]
}
