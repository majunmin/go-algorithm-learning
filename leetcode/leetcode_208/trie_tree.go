/**
  @Author: majm@ushareit.com
  @date: 2021/1/23
  @note:
**/
package leetcode_208

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		if cur.next[c-'a'] == nil {
			cur.next[c-'a'] = new(Trie)
		}
		cur = cur.next[c-'a']
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		if cur.next[c-'a'] == nil {
			return false
		}
		cur = cur.next[c-'a']
	}
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		if cur.next[c-'a'] == nil {
			return false
		}
		cur = cur.next[c-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
