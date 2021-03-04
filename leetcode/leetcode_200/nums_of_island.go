/**
  @Author: majm@ushareit.com
  @date: 2021/1/25
  @note:
**/
package leetcode_200

import (
	"container/list"
	"go-algorithm-learning/leetcode/common"
)

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

func numIslands(grid [][]byte) int {
	return solution(grid)
	//return unionset(grid)
}

// 并查集
func unionset(grid [][]byte) int {
	directions = [][]int{
		{0, 1},
		{1, 0},
	}

	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	unionFind := common.NewUnionFind(rows * cols)
	space := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				space++
				continue
			}
			for _, direction := range directions {
				newI := i + direction[0]
				newJ := j + direction[1]
				if 0 <= newI && newI < rows && 0 <= newJ && newJ < cols && grid[newI][newJ] == 1 {
					unionFind.Union(i*cols+j, newI*cols+newJ)
				}
			}
		}
	}
	return unionFind.Count() - space
}

func solution(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	count := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 {
				continue
			}
			count++
			dfs(&grid, i*col+j, col, row)
		}
	}

	return count
}

// 深度优先搜索
func dfs(grid *[][]byte, curNode int, col int, row int) {
	// terminate cond

	// process
	m, n := curNode/col, curNode%col
	(*grid)[m][n] = 0

	// 下探
	for _, direction := range directions {
		newI, newJ := m+direction[0], n+direction[1]
		if 0 <= newI && newI < row && 0 <= newJ && newJ < col && (*grid)[newI][newJ] == 1 {
			dfs(grid, newI*col+newJ, col, row)
		}
	}
}

// 广度优先搜索
func bfs(grid *[][]byte, curNode int, col int, row int) {
	queue := list.New()
	queue.PushBack(curNode)

	for queue.Len() != 0 {
		curNode := queue.Remove(queue.Front()).(int)
		m, n := curNode/col, curNode%col
		// 这里的剪枝操作必须要写,因为 bfs 遍历某一个节点时，可能被其他的路径 写成0了,不懂就debug
		if (*grid)[m][n] == 0 {
			continue
		}
		// 将当前岛屿打平
		(*grid)[m][n] = 0
		for _, direction := range directions {
			newR, newC := m+direction[0], n+direction[1]
			if newR >= 0 && newR < row && newC >= 0 && newC < col && (*grid)[newR][newC] == 1 {
				queue.PushBack(newR*col + newC)
			}
		}
	}
}
