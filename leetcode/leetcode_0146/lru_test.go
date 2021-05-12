/**
  @Author: majm@ushareit.com
  @date: 2021/2/6
  @note:
**/
package leetcode_0146

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {

	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(1, 2)
	lruCache.Put(2, 3)
	lruCache.Put(4, 1)

	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))
	fmt.Println(lruCache.Get(3))

}
