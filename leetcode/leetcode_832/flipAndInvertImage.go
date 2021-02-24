/**
  @Author: majm@ushareit.com
  @date: 2021/2/24
  @note:
**/
package leetcode_832

func flipAndInvertImage(A [][]int) [][]int {
	rows, cols := len(A), len(A[0])
	for i := 0; i < rows; i++ {
		start, end := 0, cols-1
		for end >= start {
			if end == start {
				A[i][start] = A[i][start] ^ 1
				break
			}
			A[i][start], A[i][end] = A[i][end]^1, A[i][start]^1
			start++
			end--
		}
	}
	return A
}
