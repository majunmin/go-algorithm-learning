/**
  @Author: majm@ushareit.com
  @date: 2021/2/19
  @note:
**/
package leetcode_547

import "go-algorithm-learning/leetcode/common"

func findCircleNum(isConnected [][]int) int {
	rows := len(isConnected)
	if rows <= 0 {
		return 0
	}

	uf := common.NewUnionFind(rows * rows)

	for i := 0; i < rows-1; i++ {
		for j := i + 1; j < rows; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count()
}
