/**
* Date: 2021/6/4 1:27 上午
* Author: majunmin
**/
package repeat_exercise

import "testing"

func Test_TrieTree(t *testing.T) {
	trieTree := Constructor()
	t.Log(trieTree.Search("apple"))
	trieTree.Insert("apple")
	t.Log(trieTree.Search("app"))     // false
	t.Log(trieTree.Search("apple"))   // true
	t.Log(trieTree.StartsWith("app")) // true

	trieTree.Insert("app")
	t.Log(trieTree.Search("tea"))     // true
	t.Log(trieTree.Search("apple"))   // true
	t.Log(trieTree.StartsWith("app")) // true
}
