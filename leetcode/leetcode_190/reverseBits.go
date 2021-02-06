/**
  @Author: majm@ushareit.com
  @date: 2021/2/5
  @note:
**/
package leetcode_190

//190. 颠倒二进制位
//     颠倒给定的 32 位无符号整数的二进制位
func reverseBits(num uint32) uint32 {
	res := uint32(0)
	//1. for 循环
	for i := 32; i >= 0; i-- {
		res += (num & 1) << i
		num = num >> 1
	}
	return res
}
