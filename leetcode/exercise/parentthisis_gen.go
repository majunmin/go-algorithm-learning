/**
  @Author: majm@ushareit.com
  @date: 2021/3/9
  @note:
**/
package exercise

// dp[i] 表示 n = i 时所有有效括号的组合
// dp[i] = "(dp[p]的所有有效组合)+ dp[q]的组合"
//      其中  p + q = i-1, p: 0~i-1, q: i-1~0
func generateParenthesis(n int) []string {
	dp := make(map[int][]string)
	dp[0] = []string{""}
	dp[1] = []string{"()"}

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, p := range dp[j] {
				for _, q := range dp[i-1-j] {
					dp[i] = append(dp[i], "("+p+")"+q)
				}
			}
		}
	}
	return dp[n]
}
