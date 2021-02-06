/**
  @Author: majm@ushareit.com
  @date: 2021/2/5
  @note:
**/
package leetcode_231

// 给定一个整数，编写一个函数来判断它是否是 2 的幂次方
//    ===> 2的幂次方 的数有个特点，其二进表示法 只有一个1
func isPowerOfTwo(n int) bool {

	// 位运算的 方法1    n & (n -1): 可以将 n 的最低位 清零        =>   0ms 2.2M
	return n > 0 && n&(n-1) == 0
	// 位运算的 方法2  n & (-n) : 可以将n的最低位的1 保留，其余位置0 =>   4ms 2.2M
	// return n > 0 && n & -n == n
}
