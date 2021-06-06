/**
* Date: 2021/6/2 4:44 下午
* Author: majunmin
**/
package common

import "testing"

func Test_SolveNQueues(t *testing.T) {
	t.Log(solveNQueens(4))
}

func Test_GetPath(t *testing.T) {
	t.Log(buildEmptyPath(4, 1))
	t.Log(buildEmptyPath(4, 0))
	t.Log(buildEmptyPath(4, 3))
}
