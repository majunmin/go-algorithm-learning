/**
* Date: 2021/6/6 12:54 下午
* Author: majunmin
**/
package repeat_exercise

import "testing"

func Test_NumOfIsland(t *testing.T) {
	example1 := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 1},
	}
	example2 := [][]byte{
		{1, 1, 1},
		{0, 1, 0},
		{1, 1, 1},
	}
	t.Log(numIslands(example1))
	t.Log(numIslands(example2))
	//  x + 32 = x | 32
	// x - 32  = x

	t.Log('a') // 97
	t.Log('A') // 65     'a' -'A' = 32
	t.Log('A' | 32)
	t.Log(65 ^ 32)
}
