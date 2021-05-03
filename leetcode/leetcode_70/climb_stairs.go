/**
  @Author: majm@ushareit.com
  @date: 2021/1/16
  @note:
**/
package leetcode_70

func climbStairs(n int) int {
	//return dfs(n)
	return iterForFib1(n)
}

// 1. dfs 超出时间限制
func dfs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return dfs(n) + dfs(n-1)
}

// DP
// DP 方程
// n = 1|0  :dp[n] = 1
// n > 1    :dp[n] = dp[n-1] + dp[n-2]
func iterForFib1(n int) int {
	if n == 1 {
		return 1
	}
	resArr := make([]int, n+1)
	resArr[0], resArr[1] = 1, 1
	for i := 2; i <= n; i++ {
		resArr[i] = resArr[i-1] + resArr[i-2]
	}
	return resArr[n]
}

func iterForFib2(n int) int {
	prev, next, result := 1, 1, 1
	for i := 2; i <= n; i++ {
		result = prev + next
		prev, next = next, result
	}
	return result
}
