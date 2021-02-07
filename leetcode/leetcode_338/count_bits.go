/**
  @Author: majm@ushareit.com
  @date: 2021/2/6
  @note:
**/
package leetcode_338

func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

func low1(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		result[i] = countBit(i)
	}
	return result
}

func countBit(num int) int {
	c := 0
	for num != 0 {
		c++
		num = num & (num - 1)
	}
	return c
}
