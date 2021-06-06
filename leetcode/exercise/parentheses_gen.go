/**
* Date: 2021/5/29 2:55 下午
* Author: majunmin
**/
package exercise

// 1.递归 2. 迭代  3.  dp
func genParentheses(n int) []string {
	var result []string

	if n == 0 {
		return result
	}
	m := make(map[int][]string)
	m[0] = []string{""}
	m[1] = []string{"()"}

	// 迭代递推, 自底向上, 类似  斐波那契数列
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			llist, rlist := m[j], m[i-j-1]
			for _, s1 := range llist {
				for _, s2 := range rlist {
					m[i] = append(m[i], s1+"("+s2+")")
				}
			}
		}
	}
	return m[n]
}

func genParenthesesHelper(path string, l, r, n int, result *[]string) {
	// terminate cond
	if l == n && r == n {
		*result = append(*result, path)
		return
	}
	// process +  drillDown   + revert
	if l < n {
		genParenthesesHelper(path+"(", l+1, r, n, result)
	}
	if r < l {
		genParenthesesHelper(path+")", l, r+1, n, result)
	}
}
