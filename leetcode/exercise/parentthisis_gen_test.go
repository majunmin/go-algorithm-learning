/**
  @Author: majm@ushareit.com
  @date: 2021/3/9
  @note:
**/
package exercise

import (
	"testing"
)

func TestGenParthen(t *testing.T) {
	t.Log(genParentheses(1))
	t.Log(genParentheses(2))
	t.Log(genParentheses(3))
	t.Log(genParentheses(4))
}
