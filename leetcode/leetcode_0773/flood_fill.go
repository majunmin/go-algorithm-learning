/**
* Date: 2021/6/6 10:34 上午
* Author: majunmin
**/
package leetcode_0773

// 四连通图
var direction = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

// https://leetcode-cn.com/problems/flood-fill/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	return iterSolution(image, sr, sc, newColor)
}

// bfs 广度优先搜索
func iterSolution(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	oldColor := image[sr][sc]
	// initQueue
	var queue [][]int
	queue = append(queue, []int{sr, sc})

	for i := 0; i < len(queue); i++ {
		cell := queue[i]
		curColor := image[cell[0]][cell[1]]
		if curColor != oldColor {
			continue
		}

		image[cell[0]][cell[1]] = color
		for _, d := range direction {
			newRow, newCol := cell[0]+d[0], cell[1]+d[1]
			if validPoint(newRow, newCol, len(image), len(image[0])) {
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}
	return image
}

// dfs 深度优先搜索
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)
func dfsSolution(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	dfs(&image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(result *[][]int, sr int, sc int, color, newColor int) {
	if !validPoint(sr, sc, len(*result), len((*result)[0])) {
		return
	}
	curColor := (*result)[sr][sc]
	if curColor != color || curColor == newColor {
		return
	}

	(*result)[sr][sc] = newColor

	for _, dir := range direction {
		dfs(result, sr+dir[0], sc+dir[1], color, newColor)
	}
}

// valid 坐标
func validPoint(sr int, sc int, rowLen int, colLen int) bool {
	return 0 <= sr && sr < rowLen && 0 <= sc && sc < colLen
}
