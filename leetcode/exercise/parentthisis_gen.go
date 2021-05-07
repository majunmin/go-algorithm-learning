/**
  @Author: majm@ushareit.com
  @date: 2021/3/9
  @note:
**/
package exercise

// dp[i] 表示 n = i 时所有有效括号的组合
// dp[i] = "(dp[p]的所有有效组合)+ dp[q]的组合"
//      其中  p + q = i-1, p: 0~i-1, q: i-1~0
//func generateParenthesis(n int) []string {
//	dp := make(map[int][]string)
//	dp[0] = []string{""}
//	dp[1] = []string{"()"}
//
//	for i := 2; i <= n; i++ {
//		for j := 0; j < i; j++ {
//			for _, p := range dp[j] {
//				for _, q := range dp[i-1-j] {
//					dp[i] = append(dp[i], "("+p+")"+q)
//				}
//			}
//		}
//	}
//	return dp[n]
//}

// dp方程解法 + 递归解法
func genParenthesis(n int) []string {
	// 缓存 cache
	dp := make(map[int][]string, n)
	dp[0] = []string{""}
	dp[1] = []string{"()"}

	for k := 2; k <= n; k++ {
		for i := 0; i < k; i++ {
			for _, p := range dp[i] {
				for _, q := range dp[k-i-1] {
					dp[k] = append(dp[k], "("+p+")"+q)
				}
			}
		}
	}
	return dp[n]
}

// dp方程解法 + 递归解法
func genParenthesis2(n int) []string {
	left, right := 0, 0
	var result []string
	genparenthesis("", left, right, n, &result)
	return result
}

func genparenthesis(path string, left, right, n int, result *[]string) {
	// terminate condition
	if len(path) == n*2 {
		*result = append(*result, path)
	}
	//
	if left < n {
		genparenthesis(path+"(", left+1, right, n, result)
	}
	if right < left {
		genparenthesis(path+")", left, right+1, n, result)
	}
}
