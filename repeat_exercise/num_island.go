/**
* Date: 2021/6/6 12:36 下午
* Author: majunmin
**/
package repeat_exercise

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// https://leetcode-cn.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {

	//return BFSolution(grid)
	return DFSolution(grid)
}

// 递归  dfs
func DFSolution(grid [][]byte) int {
	var numIsland int
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				numIsland++
				dfsIsland(grid, r, c)
			}
		}
	}
	return numIsland
}

func dfsIsland(grid [][]byte, r int, c int) {
	if !(0 <= r && r < len(grid) && 0 <= c && c < len(grid[r])) || grid[r][c] == 0 {
		return
	}
	// process
	grid[r][c] = 0
	// for choice in choiceList
	for _, d := range directions {
		newRow, newCol := r+d[0], c+d[1]
		dfsIsland(grid, newRow, newCol)
	}
}

func BFSolution(grid [][]byte) int {
	numOfIsland := 0
	var queue [][]int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 1 {
				continue
			}
			numOfIsland++
			queue = append(queue, []int{i, j})
			for len(queue) > 0 {
				limit := len(queue)
				for k := 0; k < limit; k++ {
					cur := queue[k]
					grid[cur[0]][cur[1]] = 0

					for _, d := range directions {
						newRow, newCol := cur[0]+d[0], cur[1]+d[1]
						if 0 <= newRow && newRow < len(grid) && 0 <= newCol && newCol < len(grid[0]) && grid[newRow][newCol] == 1 {
							queue = append(queue, []int{newRow, newCol})
						}
					}
				}
				queue = queue[limit:]
			}
		}
	}
	return numOfIsland
}
