/**
  @Author: majm@ushareit.com
  @date: 2021/1/14
  @note:
**/
package leetcode_1018

func prefixesDivBy5(A []int) []bool {

	// 1. 整个移动的过程 就类似一个 左移一位(*2) + A[i]
	// 2. 每次只对 个位数操作就可以(不用对整个 二进制数 全部转换为十进制数)
	result := make([]bool, 0, len(A))
	prefix := 0
	for i := 0; i < len(A); i++ {
		prefix = (prefix<<1 | A[i]) % 10
		result = append(result, prefix%5 == 0)
	}
	return result
}
