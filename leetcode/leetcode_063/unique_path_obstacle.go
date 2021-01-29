/**
  @Author: majm@ushareit.com
  @date: 2021/1/10
  @note:
**/
package leetcode_063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 防止开局第一个是石头
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	var dp = make([][]int, row)
	for j := 0; j < row; j++ {
		dp[j] = make([]int, col)
	}
	if obstacleGrid[0][0] != 1 { //如果起点不是障碍物，置为1，dp开始
		dp[0][0] = 1
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j] // 当前节点左侧有节点
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1] // 当前节点上侧有节点
			}
		}
	}
	return dp[row-1][col-1]
}
