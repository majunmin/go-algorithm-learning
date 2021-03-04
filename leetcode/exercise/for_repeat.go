/**
  @Author: majm@ushareit.com
  @date: 2021/2/19
  @note:
**/
package exercise

var directions = [][]int{
	{0, 1},
	{1, 0},
}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	uf := NewUnionFind(rows * cols)
	space := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				space++
				continue
			}
			for _, direction := range directions {
				newRow, newCol := i+direction[0], j+direction[1]
				if 0 <= newRow && newRow < rows && 0 <= newCol && newCol < cols && grid[newRow][newCol] == '1' {
					uf.Union(i*cols+j, newRow*cols+newCol)
				}
			}
		}
	}

	return uf.Count() - space
}

type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(count int) *UnionFind {
	if count <= 0 {
		return nil
	}

	parent := make([]int, count)
	for i := 0; i < count; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent, count: count}
}

// 返回 分组个数
func (uf UnionFind) Count() int {
	return uf.count
}

// 合并两个节点
func (uf *UnionFind) Union(p, q int) {
	p = uf.Find(p)
	q = uf.Find(q)

	if p == q {
		return
	}
	uf.parent[q] = p
	uf.count--
	return
}

// 找到顶层节点 root
func (uf UnionFind) Find(p int) int {
	for uf.parent[p] != p {
		// 这行语句是为了  压扁结构，加速Find 查找
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}
