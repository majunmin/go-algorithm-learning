/**
  @Author: majm@ushareit.com
  @date: 2021/3/8
  @note:
**/
package exercise

//状态转移方程 dp[n] = dp[n-1] + dp[n-2]
func climbStairs(n int) int {
	prev, next, result := 1, 1, 1

	for i := 2; i <= n; i++ {
		result = prev + next
		prev = next
		next = result
	}
	return result
}
