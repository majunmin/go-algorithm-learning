/**
* Date: 2021/6/4 1:10 上午
* Author: majunmin
**/
package repeat_exercise

// 只会有 小写字母
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		next: [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

	cur := this
	for i := 0; i < len(word); i++ {
		b := word[i]
		if cur.next[b-97] == nil {
			cur.next[b-97] = new(Trie)
		}
		cur = cur.next[b-97]
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	cur := this
	for i := 0; i < len(word); i++ {
		b := word[i]
		if cur.next[b-97] == nil {
			return false
		}
		cur = cur.next[b-97]
	}
	return cur.isEnd

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		b := prefix[i]
		if cur.next[b-97] == nil {
			return false
		}
		cur = cur.next[b-97]
	}
	return cur != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
