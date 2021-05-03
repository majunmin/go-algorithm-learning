/**
  @Author: majm@ushareit.com
  @date: 2021/2/5
  @note:
**/
package leetcode_0109

// 编写一个函数，输入是一个无符号整数(以二进制串的形式), 返回其二进制表达式中数字位数为 '1' 的个数(也被称为汉明重量)
func hammingWeight(num uint32) int {
	return solution(num)
}

// 利用一个巧妙的操作  x & (x - 1): 把x的 最低位的1 清零
func solution(num uint32) int {
	c := 0
	for num != 0 {
		num = num & (num - 1) // 将 num最低位 的 1 清零
		c++
	}
	return c
}

// 循环32次
func loop(num uint32) int {
	c := 0
	start := uint32(1)
	for start != 0 {
		if num&start != 0 {
			c++
		}
		start = start << 1
	}
	return c
}

// 循环32次
func loop2(num uint32) int {

	c := 0
	for num != 0 {
		if num&1 != 0 {
			c++
		}
		num = num >> 1
	}
	return c
}
