/**
  @Author: majm@ushareit.com
  @date: 2021/2/22
  @note:
**/
package leetcode_766

//  https://leetcode-cn.com/problems/toeplitz-matrix/
// 如果矩阵存储在磁盘上，并且内存有限，以至于一次最多只能将矩阵的一行加载到内存中，该怎么办？
// 如果矩阵太大，以至于一次只能将不完整的一行加载到内存中，该怎么办？
func isToeplitzMatrix(matrix [][]int) bool {
	rows, cols := len(matrix), len(matrix[0])
	if rows == 0 || cols == 0 {
		return false
	}
	if rows == 1 || cols == 1 {
		return true
	}
	return solution3(matrix)
}

// 按线读取   验证每一个对角线, 减少IO 次数
// https://leetcode-cn.com/problems/toeplitz-matrix/solution/cong-ci-pan-du-qu-cheng-ben-fen-xi-liang-f20w/
func solution3(matrix [][]int) bool {
	rows, cols := len(matrix), len(matrix[0])
	m, n := rows, cols
	for n > 0 {
		n--
		i, j := 0, n
		for val := matrix[i][j]; i < rows && j < cols; i, j = i+1, j+1 {
			if val != matrix[i][j] {
				return false
			}
		}
	}

	for m > 0 {
		m--
		i, j := m, 0
		for val := matrix[i][j]; i < rows && j < cols; i, j = i+1, j+1 {
			if val != matrix[i][j] {
				return false
			}
		}
	}
	return true
}

// 我也就能用到这种方法  -_-||
func solution2(matrix [][]int) bool {
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}

// 利用  同一捺上  row -coL = constant
func solution1(matrix [][]int) bool {
	rows, cols := len(matrix), len(matrix[0])

	cache := make(map[int]int)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			key := i - j
			val, ok := cache[key]
			if !ok {
				cache[key] = matrix[i][j]
			} else {
				if val != matrix[i][j] {
					return false
				}
			}
		}
	}
	return true
}
