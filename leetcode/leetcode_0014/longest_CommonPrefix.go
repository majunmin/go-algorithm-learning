/**
* Date: 2021/6/6 5:04 下午
* Author: majunmin
**/
package leetcode_0014

type TrieNode struct {
	End  bool
	Next [26]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func (tn *TrieNode) Insert(str string) {
	cur := tn
	for i := 0; i < len(str); i++ {
		b := str[i]
		if cur.Next[b-'a'] == nil {
			cur.Next[b-'a'] = NewTrieNode()
		}
		cur = cur.Next[b-'a']
	}
	cur.End = true
}

func (tn *TrieNode) Search(str string) bool {
	cur := tn
	for i := 0; i < len(str); i++ {
		b := str[i]
		if cur.Next[b-'a'] == nil {
			return false
		}
		cur = cur.Next[b-'a']
	}
	return cur.End
}

func (tn *TrieNode) HasPrefix(prefix string) bool {
	cur := tn
	for i := 0; i < len(prefix); i++ {
		b := prefix[i]
		if cur.Next[b-'a'] == nil {
			return false
		}
		cur = cur.Next[b-'a']
	}
	return true
}

// https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	//return solutionTrieTree(strs)
	return solution2(strs)
}

// 纵向遍历
func solution2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	length := len(strs[0])
	count := len(strs)
	for i := 0; i < length; i++ {
		c := strs[0][i]
		for j := 0; j < count; j++ {
			if len(strs[j]) == i || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func solutionTrieTree(strs []string) string {
	trieTree := NewTrieNode()
	for _, str := range strs {
		trieTree.Insert(str)
	}

	prefix := []byte{}
	cur := trieTree
	for i := 0; i < len(strs[0]); i++ {
		b := strs[0][i]
		isUnique := func(cur [26]*TrieNode) bool {
			eleNum := 0
			for i := 0; i < len(cur); i++ {
				if cur[i] != nil {
					eleNum++
				}
			}
			return eleNum == 1
		}
		if cur.Next[b-'a'] == nil || !isUnique(cur.Next) {
			break
		}
		prefix = append(prefix, b)
		cur = cur.Next[b-'a']
	}
	return string(prefix)
}
