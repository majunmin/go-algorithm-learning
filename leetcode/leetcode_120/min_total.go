/**
  @Author: majm@ushareit.com
  @date: 2021/1/16
  @note:
**/
package leetcode_120

func minimumTotal(triangle [][]int) int {
	return minTotalDP2(triangle)
}

// 3. 优化空间
func minTotalDP2(triangle [][]int) int {
	f := make([]int, len(triangle)+1)
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			f[j] = triangle[i][j] + minInt(f[j], f[j+1])
		}
	}
	return f[0]
}

// 1. 递归
// 2. DP  - 相当于最短路径和切掉了一块
//   a. 重复性(分治): problem = min(sub(i+1, j), sub(i+1,j+1)) + a[i,j]
//   b. 定义状态数组: f[i,j] = min(f(i+1, j), f(i+1,j+1)) + a[i,j]
//   c. dp方程
func minTotalDP(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = minInt(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
