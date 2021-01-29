/**
  @Author: majm@ushareit.com
  @date: 2021/1/23
  @note:
**/
package leetcode_208

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	trie := Constructor()
	trie.Insert("majm")
	trie.Insert("tea")
	trie.Insert("tead")
	trie.Insert("teamd")

	fmt.Println(trie.StartsWith("tea"))
	fmt.Println(trie.Search("tead"))
	fmt.Println(trie.Search("team"))
}
