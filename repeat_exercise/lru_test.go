/**
* Date: 2021/5/24 11:51 下午
* Author: majunmin
**/
package repeat_exercise

import (
	"testing"
)

func Test_LRU(t *testing.T) {
	lruCache := NewCache(3)
	lruCache.Add("key1", "value1")
	lruCache.Add("key2", "value2")
	lruCache.Add("key3", "value3")
	lruCache.Add("key4", "value4")
	lruCache.Add("key3", "value3-1")
	lruCache.Add("key5", "value5")
	lruCache.Add("key6", "value6")
	lruCache.Remove("key5")
	t.Log(lruCache.cache)

	t.Log(minWindow("ADOBECODEBANC", "ABC"))
	t.Log(minWindow("aa", "aa"))
}
