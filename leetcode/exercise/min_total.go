/**
  @Author: majm@ushareit.com
  @date: 2021/3/8
  @note:
**/
package exercise

func minimumTotal(triangle [][]int) int {
	return iterSolution(triangle)
}

func iterSolution(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return triangle[0][0]
	}

	for i := length - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			triangle[i][j] = min(triangle[i][j+1], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

func dfs_solution(triangle [][]int) int {
	rows := len(triangle)
	if rows == 1 {
		return triangle[0][0]
	}

	return dfs(triangle, 0, 0)
}

func dfs(triangle [][]int, row, col int) int {
	if row == len(triangle) {
		return 0
	}

	return min(dfs(triangle, row+1, col), dfs(triangle, row+1, col+1)) + triangle[row][col]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
