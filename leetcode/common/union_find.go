/**
  并查集
  @Author: majm@ushareit.com
  @date: 2021/1/25
  @note:
**/
package common

// 并查集实现
type UnionFind struct {
	// root 根节点的数量
	count  int
	parent []int
}

func NewUnionFind(count int) *UnionFind {
	u := &UnionFind{count: count}
	for i := 0; i < count; i++ {
		u.parent = append(u.parent, i)
	}
	return u
}

// 找到 顶层root节点
func (uf *UnionFind) Find(p int) int {
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}

func (uf *UnionFind) Count() int {
	return uf.count
}

// 合并两个 节点
func (uf *UnionFind) Union(p, q int) {
	rP := uf.Find(p)
	rQ := uf.Find(q)
	if rP == rQ {
		return
	}
	uf.parent[rP] = rQ
	uf.count--
}
