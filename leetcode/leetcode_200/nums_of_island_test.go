/**
  @Author: majm@ushareit.com
  @date: 2021/1/25
  @note:
**/
package leetcode_200

import (
	"fmt"
	"testing"
)

func TestNumsIsland(t *testing.T) {
	grid1 := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	grid2 := [][]byte{
		{1, 1, 1},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(numIslands(grid1))
	fmt.Println(numIslands(grid2))

}
